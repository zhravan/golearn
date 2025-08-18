package cli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"golearn/internal/exercises"
	"golearn/internal/progress"
)

func runList() error {
	items, err := exercises.List()
	if err != nil {
		return err
	}
	for _, ex := range items {
		status := "pending"
		done, _ := progress.IsCompleted(ex.Slug)
		if done {
			status = "done"
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
	items, err := exercises.List()
	if err != nil {
		return err
	}
	var anyFailed bool
	for _, ex := range items {
		if err := verifyOne(ex); err != nil {
			anyFailed = true
		}
	}
	if anyFailed {
		return errors.New("some exercises failed")
	}
	return nil
}

func verifyOne(ex exercises.Exercise) error {
	fmt.Printf("\n==> %s: %s\n", ex.Slug, ex.Title)

	cmd := exec.Command("go", "test", "-run", ex.TestRegex, "-json", "./exercises/"+ex.Slug)
	cmd.Env = append(os.Environ(), "GOFLAGS=-count=1")
	cmd.Dir = projectRoot()
	out, err := cmd.CombinedOutput()

	// Always show parsed diagnostics
	parseAndDisplayJSON(out)

	if err == nil {
		_ = progress.MarkCompleted(ex.Slug)
		fmt.Printf("PASSED %s\n", ex.Slug)
		return nil
	}
	fmt.Printf("FAILED %s\n", ex.Slug)
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
					fmt.Print(hint)
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
		fmt.Printf("%d) %s\n", i+1, h)
	}
	return nil
}

func runWatch() error {
	fmt.Println("Watch mode not yet implemented. Use your editor's file watch or rerun verify.")
	return nil
}

func runProgress() error {
	items, err := exercises.List()
	if err != nil {
		return err
	}
	doneCount := 0
	for _, ex := range items {
		done, _ := progress.IsCompleted(ex.Slug)
		if done {
			doneCount++
		}
	}
	fmt.Printf("Progress: %d/%d completed\n", doneCount, len(items))
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

func runInit() error {
	return exercises.InitAll()
}

func projectRoot() string {
	wd, _ := os.Getwd()
	return wd
}

// unused to keep future-ready; avoid linter complaints by referencing
var _ = filepath.Join
var _ = time.Now
