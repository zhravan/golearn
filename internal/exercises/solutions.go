package exercises

import (
	"embed"
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Embed the canonical solutions alongside templates. We do not expose these via CLI directly.
//
//go:embed solutions/**
var solutionsFS embed.FS

// SolutionExists reports whether we have an embedded solution for the given slug.
func SolutionExists(slug string) bool {
	root := filepath.Join("solutions", slug)
	_, err := fs.Stat(solutionsFS, root)
	return err == nil
}

// CreateSolutionSandbox creates a temporary module containing the solution implementation
// files and the original tests from templates for the given exercise slug.
// The returned directory can be used as the working directory for `go test`.
func CreateSolutionSandbox(slug string) (string, func(), error) {
	if !templateExists(slug) {
		return "", func() {}, errors.New("no template found for exercise")
	}
	if !SolutionExists(slug) {
		return "", func() {}, errors.New("no embedded solution available")
	}

	workDir, err := os.MkdirTemp("", "golearn-solution-*")
	if err != nil {
		return "", func() {}, err
	}
	cleanup := func() { _ = os.RemoveAll(workDir) }

	// 1) Create a minimal go.mod. We pin testify which is used by some tests.
	goMod := "module golearn/tmp/" + strings.ReplaceAll(slug, "_", "-") + "\n\n" +
		"go 1.22.0\n\n" +
		"require github.com/stretchr/testify v1.11.0\n"
	if err := os.WriteFile(filepath.Join(workDir, "go.mod"), []byte(goMod), 0o644); err != nil {
		cleanup()
		return "", func() {}, err
	}

	// 2) Copy tests from embedded templates
	templateRoot := filepath.Join("templates", slug)
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
		rel, err := filepath.Rel(templateRoot, path)
		if err != nil {
			return err
		}
		dest := filepath.Join(workDir, rel)
		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return err
		}
		data, err := fs.ReadFile(templatesFS, path)
		if err != nil {
			return err
		}
		return os.WriteFile(dest, data, 0o644)
	}); err != nil {
		cleanup()
		return "", func() {}, err
	}

	// 3) Copy solution implementation files (exclude *_test.go just in case)
	solutionRoot := filepath.Join("solutions", slug)
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
		rel, err := filepath.Rel(solutionRoot, path)
		if err != nil {
			return err
		}
		dest := filepath.Join(workDir, rel)
		if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
			return err
		}
		data, err := fs.ReadFile(solutionsFS, path)
		if err != nil {
			return err
		}
		return os.WriteFile(dest, data, 0o644)
	}); err != nil {
		cleanup()
		return "", func() {}, err
	}

	return workDir, cleanup, nil
}
