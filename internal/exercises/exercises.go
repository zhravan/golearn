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
		{Slug: "01_hello", Title: "Hello, Go!", TestRegex: ".*", Hints: []string{"Implement Hello() to return 'Hello, Go!'"}},
		{Slug: "02_values", Title: "Values", TestRegex: ".*", Hints: []string{"Use fmt.Sprintf to format values."}},
		{Slug: "03_variables", Title: "Variables", TestRegex: ".*", Hints: []string{"Use short declarations (:=) and return multiple values."}},
		{Slug: "04_constants", Title: "Constants", TestRegex: ".*", Hints: []string{"Use math.Pi and constant expressions."}},
		{Slug: "05_for", Title: "For", TestRegex: ".*", Hints: []string{"Accumulate a sum with a for loop."}},
		{Slug: "06_if_else", Title: "If/Else", TestRegex: ".*", Hints: []string{"Handle negative, zero, and positive cases."}},
		{Slug: "07_switch", Title: "Switch", TestRegex: ".*", Hints: []string{"Match multiple cases for weekend days."}},
		{Slug: "08_arrays", Title: "Arrays", TestRegex: ".*", Hints: []string{"Iterate with range over a fixed-size array."}},
		{Slug: "09_slices", Title: "Slices", TestRegex: ".*", Hints: []string{"Append values then compute a sum."}},
		{Slug: "10_maps", Title: "Maps", TestRegex: ".*", Hints: []string{"Use strings.Fields and map[string]int for word counts."}},
		{Slug: "11_functions", Title: "Functions", TestRegex: ".*", Hints: []string{"Pass a function and call it."}},
		{Slug: "12_multi_return", Title: "Multiple Return Values", TestRegex: ".*", Hints: []string{"Return quotient, remainder, and an error for divide-by-zero."}},
		{Slug: "13_variadic", Title: "Variadic Functions", TestRegex: ".*", Hints: []string{"Use '...' to accept any number of ints and sum them."}},
	}
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
