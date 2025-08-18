package exercises

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

//go:embed templates/**
var templatesFS embed.FS

type Exercise struct {
	Slug      string
	Title     string
	TestRegex string
	Hints     []string
}

func catalog() []Exercise {
	return []Exercise{
		{
			Slug:      "01_hello",
			Title:     "Hello, Go!",
			TestRegex: ".*",
			Hints: []string{
				"Implement Hello() to return 'Hello, Go!'",
			},
		},
	}
}

func List() ([]Exercise, error) {
	return catalog(), nil
}

func Get(slug string) (Exercise, error) {
	for _, ex := range catalog() {
		if ex.Slug == slug {
			return ex, nil
		}
	}
	return Exercise{}, fmt.Errorf("exercise not found: %s", slug)
}

func Reset(ex Exercise) error {
	// re-copy from templates over working exercises dir
	return copyExerciseTemplate(ex.Slug)
}

func InitAll() error {
	for _, ex := range catalog() {
		if err := copyExerciseTemplate(ex.Slug); err != nil {
			return err
		}
	}
	return nil
}

func copyExerciseTemplate(slug string) error {
	targetDir := filepath.Join("exercises", slug)
	// Remove and recreate to ensure a clean state
	_ = os.RemoveAll(targetDir)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}

	root := filepath.Join("templates", slug)
	return fs.WalkDir(templatesFS, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}
		dest := filepath.Join(targetDir, rel)
		if d.IsDir() {
			return os.MkdirAll(dest, 0o755)
		}
		data, err := fs.ReadFile(templatesFS, path)
		if err != nil {
			return err
		}
		return os.WriteFile(dest, data, 0o644)
	})
}

var ErrNoTemplates = errors.New("no templates found")
