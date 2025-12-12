package exercises

import (
	"embed"
	"errors"
	"io/fs"
	"os"
	pathpkg "path"
	"path/filepath"
	"strings"
)

//go:embed all:solutions/**
var solutionsFS embed.FS

// SolutionExists reports whether a solution directory exists in the embedded solutionsFS.
// dirName is the directory name under solutions/ (usually the exercise's Path).
func SolutionExists(dirName string) bool {
	root := pathpkg.Join("solutions", dirName)
	_, err := fs.Stat(solutionsFS, root)
	return err == nil
}

// CreateSolutionSandbox creates a temporary sandbox directory for the solution of the given exercise slug.
// It copies the solution implementation and test files into the sandbox.
// It returns the path to the sandbox, a cleanup function to remove it, and any error encountered.
func CreateSolutionSandbox(slug string) (string, func(), error) {
	// Look up exercise to get the correct source Directory
	ex, err := Get(slug)
	if err != nil {
		return "", func() {}, err
	}

	srcDir := ex.Path() // Uses Dir if set, else Slug

	if !templateExists(srcDir) {
		return "", func() {}, errors.New("no template found for exercise")
	}
	if !SolutionExists(srcDir) {
		return "", func() {}, errors.New("no embedded solution available")
	}

	workDir, err := os.MkdirTemp("", "golearn-solution-*")
	if err != nil {
		return "", func() {}, err
	}
	cleanup := func() { _ = os.RemoveAll(workDir) }

	// Create go.mod using the SLUG (01) for the module name
	goMod := "module golearn/tmp/" + strings.ReplaceAll(slug, "_", "-") + "\n\n" +
		"go 1.22.0\n\n" +
		"require github.com/stretchr/testify v1.11.0\n"
	if err := os.WriteFile(filepath.Join(workDir, "go.mod"), []byte(goMod), 0o644); err != nil {
		cleanup()
		return "", func() {}, err
	}

	// Copy tests from templates (using srcDir 101)
	templateRoot := filepath.Join("templates", srcDir)
	if err := fs.WalkDir(templatesFS, templateRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if !strings.HasSuffix(path, "_test.go") {
			return nil
		}

		rel, _ := filepath.Rel(templateRoot, path)
		dest := filepath.Join(workDir, rel)
		_ = os.MkdirAll(filepath.Dir(dest), 0o755)
		data, _ := fs.ReadFile(templatesFS, path)
		return os.WriteFile(dest, data, 0o644)
	}); err != nil {
		cleanup()
		return "", func() {}, err
	}

	// Copy solution implementation from solutions (using srcDir 101)
	solutionRoot := filepath.Join("solutions", srcDir)
	if err := fs.WalkDir(solutionsFS, solutionRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "_test.go") {
			return nil
		}

		rel, _ := filepath.Rel(solutionRoot, path)
		dest := filepath.Join(workDir, rel)
		_ = os.MkdirAll(filepath.Dir(dest), 0o755)
		data, _ := fs.ReadFile(solutionsFS, path)
		return os.WriteFile(dest, data, 0o644)
	}); err != nil {
		cleanup()
		return "", func() {}, err
	}

	return workDir, cleanup, nil
}
