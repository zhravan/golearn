package cli

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

// runPublish collects local progress and attempts to create a PR against the upstream repo
// adding/refreshing a JSON snapshot under contrib/leaderboard/progress/<user>.json.
// It prefers the GitHub CLI if available and the user is authenticated. Otherwise, it
// prints the snapshot and manual steps.
func runPublish(args []string) error {
	repoURL, userName, branchName, dryRun := parsePublishFlags(args)
	if repoURL == "" {
		repoURL = strings.TrimSpace(os.Getenv("GOLEARN_PUBLISH_REPO"))
		if repoURL == "" {
			repoURL = "https://github.com/zhravan/golearn"
		}
	}
	if userName == "" {
		userName = strings.TrimSpace(os.Getenv("GOLEARN_PUBLISH_USER"))
	}
	if branchName == "" {
		branchName = fmt.Sprintf("progress/%s-%s", sanitizeForFile(userName), time.Now().Format("20060102-150405"))
	}

	// Build snapshot
	snap, err := buildProgressSnapshot(userName)
	if err != nil {
		return err
	}

	if dryRun {
		b, _ := json.MarshalIndent(snap, "", "  ")
		fmt.Println(string(b))
		fmt.Println(theme.Muted("Dry-run: not creating a PR."))
		return nil
	}

	// Prefer GitHub CLI if present
	if _, lookErr := exec.LookPath("gh"); lookErr != nil {
		fmt.Println(theme.Muted("GitHub CLI not found. Printing snapshot and manual steps."))
		return printManualPublishInstructions(repoURL, snap)
	}

	// Ensure auth
	if err := exec.Command("gh", "auth", "status").Run(); err != nil {
		fmt.Println(theme.Muted("GitHub CLI not authenticated. Run 'gh auth login' first."))
		return printManualPublishInstructions(repoURL, snap)
	}

	// If username is still empty, derive from gh
	if strings.TrimSpace(userName) == "" {
		out, err := exec.Command("gh", "api", "user", "-q", ".login").Output()
		if err == nil {
			userName = strings.TrimSpace(string(out))
		}
	}

	owner, name := parseRepoOwnerAndName(repoURL)
	if owner == "" || name == "" {
		return fmt.Errorf("unable to parse repo URL: %s", repoURL)
	}

	// Ensure fork exists (no-op if already exists)
	_ = exec.Command("gh", "repo", "fork", fmt.Sprintf("%s/%s", owner, name), "--clone=false", "--remote=false").Run()

	// Work in a temporary clone of upstream
	tmpDir, err := ioutil.TempDir("", "golearn-publish-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	cloneCmd := exec.Command("git", "clone", repoURL, ".")
	cloneCmd.Dir = tmpDir
	cloneCmd.Stdout = os.Stdout
	cloneCmd.Stderr = os.Stderr
	if err := cloneCmd.Run(); err != nil {
		return fmt.Errorf("git clone failed: %w", err)
	}

	// Create branch
	if err := runCmd(tmpDir, "git", "checkout", "-b", branchName); err != nil {
		return err
	}

	// Write snapshot file
	progressDir := filepath.Join(tmpDir, "contrib", "leaderboard", "progress")
	if err := os.MkdirAll(progressDir, 0o755); err != nil {
		return err
	}
	fileName := sanitizeForFile(userName)
	if fileName == "" {
		fileName = "anonymous"
	}
	filePath := filepath.Join(progressDir, fileName+".json")
	b, _ := json.MarshalIndent(snap, "", "  ")
	if err := os.WriteFile(filePath, b, 0o644); err != nil {
		return err
	}

	// Update README leaderboard section
	if err := updateLeaderboardReadme(tmpDir, progressDir); err != nil {
		fmt.Printf("Warning: could not update README leaderboard: %v\n", err)
	}

	if err := runCmd(tmpDir, "git", "add", filePath); err != nil {
		return err
	}
	_ = runCmd(tmpDir, "git", "add", filepath.Join(tmpDir, "README.md"))
	msg := fmt.Sprintf("chore(leaderboard): add progress for %s (%d/%d)", snap.User, snap.CompletedCount, snap.TotalCount)
	if err := runCmd(tmpDir, "git", "commit", "-m", msg); err != nil {
		return err
	}

	// Add fork remote and push
	forkURL := fmt.Sprintf("https://github.com/%s/%s.git", userName, name)
	_ = runCmd(tmpDir, "git", "remote", "remove", "fork")
	if err := runCmd(tmpDir, "git", "remote", "add", "fork", forkURL); err != nil {
		return err
	}
	if err := runCmd(tmpDir, "git", "push", "-u", "fork", branchName+":"+branchName); err != nil {
		return err
	}

	// Create PR against upstream
	prTitle := fmt.Sprintf("Add progress for %s (%d/%d)", snap.User, snap.CompletedCount, snap.TotalCount)
	prBody := "Automated progress publish from golearn CLI. This adds/updates your progress snapshot for the leaderboard."
	prCmd := exec.Command("gh", "pr", "create",
		"--repo", fmt.Sprintf("%s/%s", owner, name),
		"--head", fmt.Sprintf("%s:%s", userName, branchName),
		"--base", "main",
		"--title", prTitle,
		"--body", prBody,
	)
	prCmd.Dir = tmpDir
	prCmd.Stdout = os.Stdout
	prCmd.Stderr = os.Stderr
	if err := prCmd.Run(); err != nil {
		fmt.Println(theme.Muted("Could not auto-create PR. You may need to open it manually."))
		fmt.Printf("Branch pushed to %s:%s\n", forkURL, branchName)
		fmt.Printf("Open a PR against %s/%s with head %s:%s\n", owner, name, userName, branchName)
		return nil
	}

	return nil
}

func parsePublishFlags(args []string) (repo, user, branch string, dry bool) {
	for _, a := range args {
		if strings.HasPrefix(a, "--repo=") {
			repo = strings.TrimSpace(strings.TrimPrefix(a, "--repo="))
			continue
		}
		if strings.HasPrefix(a, "--user=") {
			user = strings.TrimSpace(strings.TrimPrefix(a, "--user="))
			continue
		}
		if strings.HasPrefix(a, "--branch=") {
			branch = strings.TrimSpace(strings.TrimPrefix(a, "--branch="))
			continue
		}
		if a == "--dry-run" || a == "--dry" {
			dry = true
			continue
		}
	}
	return
}

type progressSnapshot struct {
	User           string   `json:"user"`
	CompletedCount int      `json:"completed_count"`
	TotalCount     int      `json:"total_count"`
	Percent        int      `json:"percent"`
	CompletedSlugs []string `json:"completed_slugs"`
	Timestamp      string   `json:"timestamp"`
}

func buildProgressSnapshot(user string) (progressSnapshot, error) {
	catalog, err := exercises.ListAll()
	if err != nil {
		return progressSnapshot{}, err
	}
	var all []exercises.Exercise
	all = append(all, catalog.Concepts...)
	all = append(all, catalog.Projects...)
	sort.Slice(all, func(i, j int) bool { return all[i].Slug < all[j].Slug })

	var completed []string
	for _, ex := range all {
		ok, _ := progress.IsCompleted(ex.Slug)
		if ok {
			completed = append(completed, ex.Slug)
		}
	}
	percent := 0
	if len(all) > 0 {
		percent = int(float64(len(completed)) / float64(len(all)) * 100)
	}
	return progressSnapshot{
		User:           user,
		CompletedCount: len(completed),
		TotalCount:     len(all),
		Percent:        percent,
		CompletedSlugs: completed,
		Timestamp:      time.Now().UTC().Format(time.RFC3339),
	}, nil
}

func sanitizeForFile(s string) string {
	s = strings.TrimSpace(strings.ToLower(s))
	if s == "" {
		return ""
	}
	// replace non-alphanumeric with '-'
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			b.WriteRune(r)
		} else {
			b.WriteRune('-')
		}
	}
	return strings.Trim(b.String(), "-")
}

func parseRepoOwnerAndName(repoURL string) (owner, name string) {
	url := strings.TrimSpace(repoURL)
	url = strings.TrimSuffix(url, ".git")
	if strings.Contains(url, "github.com") {
		// handle https and ssh
		if strings.HasPrefix(url, "git@") {
			// git@github.com:owner/name(.git)
			parts := strings.SplitN(url, ":", 2)
			if len(parts) == 2 {
				rest := parts[1]
				segs := strings.Split(strings.TrimPrefix(rest, "/"), "/")
				if len(segs) >= 2 {
					return segs[0], segs[1]
				}
			}
		} else {
			// https://github.com/owner/name
			idx := strings.Index(url, "github.com/")
			if idx >= 0 {
				rest := url[idx+len("github.com/"):]
				segs := strings.Split(strings.TrimPrefix(rest, "/"), "/")
				if len(segs) >= 2 {
					return segs[0], segs[1]
				}
			}
		}
	}
	return "", ""
}

func runCmd(dir string, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func printManualPublishInstructions(repoURL string, snap progressSnapshot) error {
	b, _ := json.MarshalIndent(snap, "", "  ")
	fmt.Println(string(b))
	fmt.Println()
	fmt.Println(theme.Muted("Manual steps to publish your progress:"))
	fmt.Println("1) Fork the repository if not already: https://github.com/zhravan/golearn")
	fmt.Println("2) Clone upstream, create a branch, add the JSON under contrib/leaderboard/progress/<your-username>.json, commit, push to your fork, and open a PR.")
	fmt.Println("   Repo:", repoURL)
	return nil
}

type leaderRow struct {
	User      string
	Timestamp string
}

func updateLeaderboardReadme(repoDir string, progressDir string) error {
	entries, err := os.ReadDir(progressDir)
	if err != nil {
		return err
	}
	var rows []leaderRow
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(progressDir, e.Name()))
		if err != nil {
			continue
		}
		var ps progressSnapshot
		if err := json.Unmarshal(b, &ps); err != nil {
			continue
		}
		if ps.TotalCount > 0 && ps.CompletedCount == ps.TotalCount {
			rows = append(rows, leaderRow{User: ps.User, Timestamp: ps.Timestamp})
		}
	}
	// sort by timestamp ascending, then username
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].Timestamp == rows[j].Timestamp {
			return rows[i].User < rows[j].User
		}
		return rows[i].Timestamp < rows[j].Timestamp
	})

    // Build dynamic content only within markers
    var section strings.Builder
    section.WriteString("<!-- START_LEADERBOARD -->\n")
    if len(rows) == 0 {
        section.WriteString("No completions yet. Be the first!\n")
    } else {
        section.WriteString("| Image | Username | Date |\n")
        section.WriteString("|---|---|---|\n")
        for _, r := range rows {
            avatar := fmt.Sprintf("https://github.com/%s.png?size=64", r.User)
            profile := fmt.Sprintf("https://github.com/%s", r.User)
            section.WriteString(fmt.Sprintf("| ![%s](%s) | [%s](%s) | %s |\n", r.User, avatar, r.User, profile, r.Timestamp))
        }
    }
    section.WriteString("<!-- END_LEADERBOARD -->\n")

    readmePath := filepath.Join(repoDir, "README.md")
    old, err := os.ReadFile(readmePath)
    if err != nil {
        // If README missing, create a new README with header and section
        var full strings.Builder
        full.WriteString("## Leaderboard\n\n")
        full.WriteString("The following users have completed all exercises (ascending by completion time):\n\n")
        full.WriteString(section.String())
        return os.WriteFile(readmePath, []byte(full.String()), 0o644)
    }

    // If markers are present, replace content between them; else append full section at end
    updated := replaceBetweenMarkers(string(old), "<!-- START_LEADERBOARD -->", "<!-- END_LEADERBOARD -->", section.String())
    if updated == string(old) {
        // markers not found; append header + section
        var full strings.Builder
        full.WriteString(string(old))
        if !strings.HasSuffix(string(old), "\n") {
            full.WriteString("\n")
        }
        full.WriteString("\n## Leaderboard\n\n")
        full.WriteString("The following users have completed all exercises (ascending by completion time):\n\n")
        full.WriteString(section.String())
        updated = full.String()
    }
    return os.WriteFile(readmePath, []byte(updated), 0o644)
}

func replaceBetweenMarkers(orig string, startMarker string, endMarker string, replacement string) string {
	startIdx := strings.Index(orig, startMarker)
	endIdx := strings.Index(orig, endMarker)
	if startIdx == -1 || endIdx == -1 || endIdx < startIdx {
		// append section at end
		if strings.HasSuffix(orig, "\n") {
			return orig + "\n" + replacement
		}
		return orig + "\n\n" + replacement
	}
	endIdx += len(endMarker)
	return orig[:startIdx] + replacement + orig[endIdx:]
}
