package exercises

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

//go:embed templates/**
var templatesFS embed.FS

//go:embed Catalog/**
var catalogFS embed.FS

type Exercise struct {
	Slug      string   `yaml:"slug"`
	Dir       string   `yaml:"dir,omitempty"` // Maps to folder name if different from slug
	Title     string   `yaml:"title"`
	TestRegex string   `yaml:"test_regex"`
	Hints     []string `yaml:"hints"`
}

// Path returns the directory name in templates/solutions.
func (e Exercise) Path() string {
	if e.Dir != "" {
		return e.Dir
	}
	return e.Slug
}

type Catalog struct {
	Concepts []Exercise `yaml:"concepts"`
	Projects []Exercise `yaml:"projects"`
}

var (
	catalogMu   sync.Mutex
	catalogData *Catalog

	catalogLoader = func() (Catalog, error) {
		return loadCatalogFromFS(catalogFS)
	}
)

// catalog returns the singleton Catalog instance.
func catalog() Catalog {
	catalogMu.Lock()
	defer catalogMu.Unlock()

	if catalogData != nil {
		return *catalogData
	}

	c, err := catalogLoader()
	if err != nil {
		// If loading fails, use fallback
		f := fallbackCatalog()
		catalogData = &f
		return *catalogData
	}

	catalogData = &c
	return *catalogData
}

// loadCatalogFromFS scans the provided filesystem for exercise definitions.
// It accepts fs.FS to allow testing with fstest.MapFS.
func loadCatalogFromFS(fsys fs.FS) (Catalog, error) {
	concepts, err := loadExercisesDir(fsys, "Catalog/Concepts")
	if err != nil {
		return Catalog{}, err
	}
	projects, err := loadExercisesDir(fsys, "Catalog/Projects")
	if err != nil {
		return Catalog{}, err
	}
	if len(concepts) == 0 && len(projects) == 0 {
		return Catalog{}, errors.New("no exercises found in Catalog/")
	}
	seen := map[string]string{} // key -> origin
	check := func(origin string, ex Exercise) error {
		for _, k := range []string{ex.Slug, ex.Dir} {
			if k == "" {
				continue
			}
			if prev, ok := seen[k]; ok {
				return fmt.Errorf("duplicate exercise identifier %q: %s vs %s", k, prev, origin)
			}
			seen[k] = origin
		}
		return nil
	}
	for _, ex := range concepts {
		if err := check("Catalog/Concepts/"+ex.Slug, ex); err != nil {
			return Catalog{}, err
		}
	}
	for _, ex := range projects {
		if err := check("Catalog/Projects/"+ex.Slug, ex); err != nil {
			return Catalog{}, err
		}
	}

	return Catalog{Concepts: concepts, Projects: projects}, nil
}

// loadExercisesDir reads all .yaml/.yml files in a directory and returns sorted exercises.
func loadExercisesDir(fsys fs.FS, dir string) ([]Exercise, error) {
	var exercises []Exercise
	entries, err := fs.ReadDir(fsys, dir)
	if errors.Is(err, fs.ErrNotExist) {
		return []Exercise{}, nil
	}
	if err != nil {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || (!strings.HasSuffix(e.Name(), ".yaml") && !strings.HasSuffix(e.Name(), ".yml")) {
			continue
		}
		data, err := fs.ReadFile(fsys, path.Join(dir, e.Name()))
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", e.Name(), err)
		}
		var ex Exercise
		if err := yaml.Unmarshal(data, &ex); err != nil {
			return nil, fmt.Errorf("parse %s: %w", e.Name(), err)
		}
		if ex.Slug == "" {
			return nil, fmt.Errorf("exercise in %s missing slug", e.Name())
		}
		// Hardening: keep identifiers as simple directory names.
		for fieldName, v := range map[string]string{"slug": ex.Slug, "dir": ex.Dir} {
			if v == "" {
				continue
			}
			if strings.Contains(v, "/") || strings.Contains(v, `\`) || strings.Contains(v, "..") || path.Clean(v) != v {
				return nil, fmt.Errorf("%s in %s is not a safe directory name: %q", fieldName, e.Name(), v)
			}
		}
		exercises = append(exercises, ex)
	}
	sort.Slice(exercises, func(i, j int) bool { return exercises[i].Slug < exercises[j].Slug })
	return exercises, nil
}

// fallbackCatalog provides a minimal hardcoded catalog if loading fails.
func fallbackCatalog() Catalog {
	return Catalog{
		Concepts: []Exercise{{
			Slug:      "01_hello",
			Title:     "Hello, Go!",
			TestRegex: ".*",
			Hints:     []string{"Implement Hello() to return 'Hello, Go!'"},
		}},
	}
}

// discoverLocal checks the local exercises/ directory for any user-added exercises.
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
			items = append(items, Exercise{Slug: slug, Title: slug, TestRegex: ".*"})
		}
	}
	return items, nil
}

// ListAll returns all exercises from the catalog or local directory if present.
func ListAll() (Catalog, error) {
	locals, err := discoverLocal()
	if err != nil {
		return Catalog{}, err
	}
	if len(locals) > 0 {
		return Catalog{Concepts: locals}, nil
	}
	return catalog(), nil
}

// Get looks up an exercise by its slug or directory name.
func Get(slug string) (Exercise, error) {
	for _, ex := range catalog().Concepts {
		if ex.Slug == slug || ex.Dir == slug {
			return ex, nil
		}
	}
	for _, ex := range catalog().Projects {
		if ex.Slug == slug || ex.Dir == slug {
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

// Reset restores the exercise files from the embedded template.
func Reset(ex Exercise) error {
	if !templateExists(ex.Path()) {
		return fmt.Errorf("reset unsupported for non-embedded exercise '%s'", ex.Slug)
	}
	return copyExerciseTemplate(ex)
}

// templateExists reports whether a template directory exists in the embedded templatesFS.
func templateExists(dirName string) bool {
	root := path.Join("templates", dirName)
	_, err := fs.Stat(templatesFS, root)
	return err == nil
}

// InitAll initializes all exercises from the embedded templates.
func InitAll() error {
	for _, ex := range catalog().Concepts {
		if err := copyExerciseTemplate(ex); err != nil {
			return err
		}
	}
	for _, ex := range catalog().Projects {
		if err := copyExerciseTemplate(ex); err != nil {
			return err
		}
	}
	return nil
}

// copyExerciseTemplate copies the template files for an exercise into the local exercises/ directory.
func copyExerciseTemplate(ex Exercise) error {
	targetDir := filepath.Join("exercises", ex.Slug)
	_ = os.RemoveAll(targetDir)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return err
	}

	root := path.Join("templates", ex.Path())
	return fs.WalkDir(templatesFS, root, func(fpath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(root, filepath.FromSlash(fpath))
		if err != nil {
			rel = strings.TrimPrefix(fpath, root+"/")
		}
		if rel == "." {
			return nil
		}
		dest := filepath.Join(targetDir, filepath.FromSlash(rel))
		if d.IsDir() {
			return os.MkdirAll(dest, 0o755)
		}
		data, err := fs.ReadFile(templatesFS, fpath)
		if err != nil {
			return err
		}
		return os.WriteFile(dest, data, 0o644)
	})
}

var ErrNoTemplates = errors.New("no templates found")
