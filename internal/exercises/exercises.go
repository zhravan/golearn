package exercises

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v3"
)

//go:embed templates/**
var templatesFS embed.FS

//go:embed catalog.yaml
var catalogFS embed.FS

type Exercise struct {
	Slug      string   `yaml:"slug"`
	Title     string   `yaml:"title"`
	TestRegex string   `yaml:"test_regex"`
	Hints     []string `yaml:"hints"`
}

var (
	catalogOnce sync.Once
	catalogData []Exercise
)

func catalog() []Exercise {
	catalogOnce.Do(func() {
		b, err := catalogFS.ReadFile("catalog.yaml")
		if err != nil {
			// Fallback minimal catalog
			catalogData = []Exercise{{
				Slug:      "01_hello",
				Title:     "Hello, Go!",
				TestRegex: ".*",
				Hints:     []string{"Implement Hello() to return 'Hello, Go!'"},
			}}
			return
		}
		var items []Exercise
		if err := yaml.Unmarshal(b, &items); err != nil || len(items) == 0 {
			catalogData = []Exercise{{
				Slug:      "01_hello",
				Title:     "Hello, Go!",
				TestRegex: ".*",
				Hints:     []string{"Implement Hello() to return 'Hello, Go!'"},
			}}
			return
		}
		catalogData = items
	})
	return catalogData
}

func discoverLocal() ([]Exercise, error) {
	var items []Exercise
	entries, err := os.ReadDir("exercises")
	if errors.Is(err, os.ErrNotExist) {
		return items, nil
	}
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() {
			slug := e.Name()
			items = append(items, Exercise{
				Slug:      slug,
				Title:     slug,
				TestRegex: ".*",
				Hints:     nil,
			})
		}
	}
	return items, nil
}

func List() ([]Exercise, error) {
	locals, err := discoverLocal()
	if err != nil {
		return nil, err
	}
	if len(locals) > 0 {
		return locals, nil
	}
	return catalog(), nil
}

func Get(slug string) (Exercise, error) {
	for _, ex := range catalog() {
		if ex.Slug == slug {
			return ex, nil
		}
	}
	locals, err := discoverLocal()
	if err != nil {
		return Exercise{}, err
	}
	for _, ex := range locals {
		if ex.Slug == slug {
			return ex, nil
		}
	}
	return Exercise{}, fmt.Errorf("exercise not found: %s", slug)
}

func Reset(ex Exercise) error {
	// Only supported for built-in embedded templates
	if !templateExists(ex.Slug) {
		return fmt.Errorf("reset unsupported for non-embedded exercise '%s'", ex.Slug)
	}
	return copyExerciseTemplate(ex.Slug)
}

func templateExists(slug string) bool {
	root := filepath.Join("templates", slug)
	_, err := fs.Stat(templatesFS, root)
	return err == nil
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
