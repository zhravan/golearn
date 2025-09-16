package cli

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/zhravan/golearn/internal/cli/theme"
	"github.com/zhravan/golearn/internal/exercises"
	"github.com/zhravan/golearn/internal/progress"
)

func runList() error {
	cat, err := exercises.ListAll()
	if err != nil {
		return err
	}

	if h := theme.Heading("Concepts"); h != "" {
		fmt.Println(h)
	}
	for _, ex := range cat.Concepts {
		status := "pending"
		done, _ := progress.IsCompleted(ex.Slug)
		if done {
			status = theme.Success("done")
		}
		if !done {
			status = theme.Muted(status)
		}
		fmt.Printf("%s - %s [%s]\n", ex.Slug, ex.Title, status)
	}

	if h := theme.Heading("Projects"); h != "" {
		fmt.Println(h)
	}
	for _, ex := range cat.Projects {
		status := "pending"
		done, _ := progress.IsCompleted(ex.Slug)
		if done {
			status = theme.Success("done")
		}
		if !done {
			status = theme.Muted(status)
		}
		fmt.Printf("%s - %s [%s]\n", ex.Slug, ex.Title, status)
	}

	return nil
}

func runVerify(name string) error {
	if name != "" {
		ex, err := exercises.Get(name)
		if err != nil {
			return err
		}
		return verifyOne(ex)
	}

	cat, err := exercises.ListAll()
	if err != nil {
		return err
	}

	var allExercises []exercises.Exercise
	allExercises = append(allExercises, cat.Concepts...)
	allExercises = append(allExercises, cat.Projects...)

	var anyFailed bool
	for _, ex := range allExercises {
		if err := verifyOne(ex); err != nil {
			anyFailed = true
		}
	}
	if anyFailed {
		return errors.New("some exercises failed")
	}
	return nil
}

// runVerifyWithOptions extends verification to support running against the embedded
// solution implementation. When useSolution is true, name must be provided.
func runVerifyWithOptions(name string, useSolution bool) error {
	if !useSolution {
		return runVerify(name)
	}
	if strings.TrimSpace(name) == "" {
		return errors.New("--solution requires a specific exercise name")
	}
	ex, err := exercises.Get(name)
	if err != nil {
		return err
	}

	fmt.Printf("\n%s\n", theme.Heading("==> "+ex.Slug+": "+ex.Title+" (solution)"))

	dir, cleanup, err := exercises.CreateSolutionSandbox(ex.Slug)
	if err != nil {
		return err
	}
	defer cleanup()

	cmd := exec.Command("go", "test", "-run", ex.TestRegex, "-json", ".")
	cmd.Env = append(os.Environ(), "GOFLAGS=-count=1")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()

	parseAndDisplayJSON(out)

	if err == nil {
		// Do NOT mark progress when validating with solutions
		fmt.Printf("%s %s (solution)\n", theme.Success("PASSED"), ex.Slug)
		return nil
	}
	fmt.Printf("%s %s (solution)\n", theme.Error("FAILED"), ex.Slug)
	return err
}

func verifyOne(ex exercises.Exercise) error {
	fmt.Printf("\n%s\n", theme.Heading("==> "+ex.Slug+": "+ex.Title))

	cmd := exec.Command("go", "test", "-run", ex.TestRegex, "-json", "./exercises/"+ex.Slug)
	cmd.Env = append(os.Environ(), "GOFLAGS=-count=1")
	cmd.Dir = projectRoot()
	out, err := cmd.CombinedOutput()

	// Always show parsed diagnostics
	parseAndDisplayJSON(out)

	if err == nil {
		_ = progress.MarkCompleted(ex.Slug)
		fmt.Printf("%s %s\n", theme.Success("PASSED"), ex.Slug)
		return nil
	}
	fmt.Printf("%s %s\n", theme.Error("FAILED"), ex.Slug)
	return err
}

func parseAndDisplayJSON(out []byte) {
	dec := json.NewDecoder(bytes.NewReader(out))
	for {
		var ev map[string]any
		if err := dec.Decode(&ev); err != nil {
			break
		}
		action, _ := ev["Action"].(string)
		if action == "output" {
			if line, ok := ev["Output"].(string); ok {
				// Surface test failure/output lines
				if strings.Contains(line, "--- FAIL") || strings.Contains(line, "FAIL\t") || strings.Contains(line, "Error:") || strings.Contains(line, "Hello() =") {
					fmt.Print(line)
				}
				if hint := hintFromCompiler(line); hint != "" {
					fmt.Print(theme.Hint(hint))
				}
			}
		}
	}
}

func hintFromCompiler(line string) string {
	switch {
	case strings.Contains(line, "undefined:"):
		return "Hint: You may need to implement the missing symbol or fix the import path.\n"
	case strings.Contains(line, "cannot use") && strings.Contains(line, "(type") && strings.Contains(line, "as type"):
		return "Hint: Type mismatch. Check function signatures and types used in assertions.\n"
	default:
		return ""
	}
}

func runHint(name string) error {
	ex, err := exercises.Get(name)
	if err != nil {
		return err
	}
	for i, h := range ex.Hints {
		fmt.Printf("%d) %s\n", i+1, theme.Hint(h))
	}
	return nil
}

// runSolution implements the hint-first flow for solutions. It never prints
// solution code in the CLI; instead it offers hints or a GitHub link.
func runSolution(name string) error {
	ex, err := exercises.Get(name)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", theme.Emph("Are you sure you want to view the solution?"))
	fmt.Print("Why not take a hint first? View hints now? [y/N]: ")

	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	line = strings.TrimSpace(strings.ToLower(line))
	if line == "y" || line == "yes" {
		return runHint(ex.Slug)
	}

	branch := strings.TrimSpace(os.Getenv("GOLEARN_SOLUTIONS_BRANCH"))
	if branch == "" {
		branch = "main"
	}
	link := fmt.Sprintf("https://github.com/zhravan/golearn/tree/%s/internal/exercises/solutions/%s", branch, ex.Slug)
	fmt.Printf("View solution on GitHub: %s\n", link)
	fmt.Printf("%s\n", theme.Muted("Tip: run 'golearn verify "+ex.Slug+" --solution' to validate the solution against the tests."))
	return nil
}

func runWatch() error {
	root := filepath.Join(projectRoot(), "exercises")
	if _, err := os.Stat(root); errors.Is(err, os.ErrNotExist) {
		fmt.Println("No exercises directory found. Run 'golearn init' first or ensure exercises are present.")
		return nil
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	if err := addWatchDirs(watcher, root); err != nil {
		return err
	}

	fmt.Println(theme.Muted("Watching for changes. Press Ctrl+C to stop."))

	// Debounce verifications per slug
	timers := map[string]*time.Timer{}
	resetTimer := func(slug string) {
		if slug == "" {
			return
		}
		if t, ok := timers[slug]; ok {
			if !t.Stop() {
				select {
				case <-t.C:
				default:
				}
			}
		}
		timers[slug] = time.AfterFunc(300*time.Millisecond, func() {
			ex, err := exercises.Get(slug)
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			_ = verifyOne(ex)
		})
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	for {
		select {
		case ev, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if ev.Op&(fsnotify.Create) != 0 {
				fi, err := os.Stat(ev.Name)
				if err == nil && fi.IsDir() {
					_ = addWatchDirs(watcher, ev.Name)
				}
			}
			if slug := slugFromPath(ev.Name); slug != "" {
				resetTimer(slug)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			fmt.Printf("%s %v\n", theme.Error("watch error:"), err)
		case <-stop:
			fmt.Println(theme.Muted("Stopping watch mode."))
			return nil
		}
	}
}

func runProgress() error {
	cat, err := exercises.ListAll()
	if err != nil {
		return err
	}

	var allExercises []exercises.Exercise
	allExercises = append(allExercises, cat.Concepts...)
	allExercises = append(allExercises, cat.Projects...)

	sort.Slice(allExercises, func(i, j int) bool { return allExercises[i].Slug < allExercises[j].Slug })
	doneCount := 0
	statuses := make([]bool, len(allExercises))
	for i, ex := range allExercises {
		done, _ := progress.IsCompleted(ex.Slug)
		statuses[i] = done
		if done {
			doneCount++
		}
	}

	// Clear screen only when appropriate and render dashboard
	if theme.MaybeClearScreen() {
		// cleared
	}
	fmt.Println(theme.Heading("GoLearn Progress Dashboard"))
	fmt.Println(strings.Repeat("=", 26))
	width := progressBarWidth()
	fmt.Printf("\nCompleted: %d/%d\n", doneCount, len(allExercises))
	fmt.Println(renderProgressBar(doneCount, len(allExercises), width))
	fmt.Println()
	fmt.Println(theme.Emph("Concepts:"))
	for _, ex := range cat.Concepts {
		box := "[ ]"
		done, _ := progress.IsCompleted(ex.Slug)
		if done {
			box = theme.Success("[x]")
		}
		fmt.Printf(" %s %s - %s\n", box, ex.Slug, ex.Title)
	}
	fmt.Println()
	fmt.Println(theme.Emph("Projects:"))
	for _, ex := range cat.Projects {
		box := "[ ]"
		done, _ := progress.IsCompleted(ex.Slug)
		if done {
			box = theme.Success("[x]")
		}
		fmt.Printf(" %s %s - %s\n", box, ex.Slug, ex.Title)
	}
	fmt.Println()
	fmt.Println(theme.Muted("Tip: run 'golearn verify <slug>' to test an exercise or 'golearn watch' for auto-verify."))
	return nil
}

func runReset(name string) error {
	if name == "" {
		return errors.New("reset requires an exercise name")
	}
	ex, err := exercises.Get(name)
	if err != nil {
		return err
	}
	return exercises.Reset(ex)
}

func runInit(repo, dir string) error {
	if strings.TrimSpace(repo) == "" {
		return exercises.InitAll()
	}

	targetDir := strings.TrimSpace(dir)
	if targetDir == "" {
		// derive folder from repo, e.g. https://github.com/org/repo(.git) -> repo
		repoPath := repo
		if idx := strings.LastIndex(repoPath, "/"); idx >= 0 && idx < len(repoPath)-1 {
			repoPath = repoPath[idx+1:]
		}
		targetDir = strings.TrimSuffix(repoPath, ".git")
		if targetDir == "" {
			targetDir = "golearn-exercises"
		}
	}

	cmd := exec.Command("git", "clone", repo, targetDir)
	cmd.Dir = projectRoot()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone failed: %w", err)
	}

	// Ensure local progress dir exists in the new workspace
	_ = os.MkdirAll(filepath.Join(targetDir, ".golearn"), 0o755)
	fmt.Printf("Cloned %s into %s\n", repo, targetDir)
	return nil
}

func projectRoot() string {
	wd, _ := os.Getwd()
	return wd
}

// unused to keep future-ready; avoid linter complaints by referencing
var _ = filepath.Join
var _ = time.Now

// addWatchDirs walks the provided directory and registers all subdirectories with the watcher.
func addWatchDirs(w *fsnotify.Watcher, dir string) error {
	return filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return w.Add(path)
		}
		return nil
	})
}

// slugFromPath extracts the exercise slug from a path like exercises/<slug>/... .
func slugFromPath(path string) string {
	path = filepath.ToSlash(path)
	idx := strings.Index(path, "exercises/")
	if idx == -1 {
		return ""
	}
	rest := path[idx+len("exercises/"):]
	parts := strings.Split(rest, "/")
	if len(parts) == 0 || parts[0] == "" {
		return ""
	}
	return parts[0]
}

func renderProgressBar(completed, total, width int) string {
	if total <= 0 {
		total = 1
	}
	if width <= 0 {
		width = 40
	}
	filled := int(float64(completed) / float64(total) * float64(width))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	bar := strings.Repeat("#", filled) + strings.Repeat("-", width-filled)
	percent := int(float64(completed) / float64(total) * 100)
	return fmt.Sprintf("[%s] %d%%", bar, percent)
}

func progressBarWidth() int {
	// Try COLUMNS env var; fallback to 60, leave margin for text
	if v := os.Getenv("COLUMNS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 10 {
			w := n - 10
			if w < 10 {
				w = 10
			}
			if w > 80 {
				w = 80
			}
			return w
		}
	}
	return 60
}
