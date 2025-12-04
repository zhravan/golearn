package exercises

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"gopkg.in/yaml.v3"
)

//go:embed templates/**
var templatesFS embed.FS

//go:embed Catalog/**
var catalogFS embed.FS

type Exercise struct {
	Slug      string   `yaml:"slug"`
	Title     string   `yaml:"title"`
	TestRegex string   `yaml:"test_regex"`
	Hints     []string `yaml:"hints"`
}

type Catalog struct {
	Concepts []Exercise
	Projects []Exercise
}

// --- Catalog Loader Infrastructure ---

// defaultCatalogLoader loads the catalog from embedded FS.
// Tests override this to inject fake catalogs.
var defaultCatalogLoader = func() (Catalog, error) {
	return loadCatalogFromFS(catalogFS)
}

var (
	catalogMu   sync.Mutex
	catalogOnce sync.Once
	catalogData Catalog
)

// withTestCatalogLoader temporarily overrides the catalog loader
// and resets internal singleton state for the duration of the test.
func withTestCatalogLoader(loader func() (Catalog, error), fn func()) {
	catalogMu.Lock()

	oldLoader := defaultCatalogLoader

	// override loader + reset singleton
	defaultCatalogLoader = loader
	catalogOnce = sync.Once{}
	catalogData = Catalog{}

	catalogMu.Unlock()

	fn()

	// restore loader, and reset the once/data again
	catalogMu.Lock()
	defaultCatalogLoader = oldLoader
	catalogOnce = sync.Once{}
	catalogData = Catalog{}
	catalogMu.Unlock()
}

// Get the singleton catalog instance.
// Loads from embedded FS on first call,
// or falls back to default if loading fails.
func catalog() Catalog {
	catalogOnce.Do(func() {
		cat, err := defaultCatalogLoader()
		if err != nil || (len(cat.Concepts) == 0 && len(cat.Projects) == 0) {
			catalogData = fallbackCatalog()
			return
		}
		catalogData = cat
	})
	return catalogData
}

// Load catalog from a given FS
// Returns a Catalog struct.
func loadCatalogFromFS(fsys fs.FS) (Catalog, error) {
	concepts, err := loadExercisesDir(fsys, "Catalog/Concepts")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return Catalog{}, err
	}

	projects, err := loadExercisesDir(fsys, "Catalog/Projects")
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return Catalog{}, err
	}

	// Deterministic ordering
	sort.Slice(concepts, func(i, j int) bool { return concepts[i].Slug < concepts[j].Slug })
	sort.Slice(projects, func(i, j int) bool { return projects[i].Slug < projects[j].Slug })

	return Catalog{
		Concepts: concepts,
		Projects: projects,
	}, nil
}

// Load all exercises from a given directory in the FS
// Returns a slice of Exercise structs.
func loadExercisesDir(fsys fs.FS, dir string) ([]Exercise, error) {
	entries, err := fs.ReadDir(fsys, dir)
	if err != nil {
		return nil, err
	}

	var out []Exercise

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		name := e.Name()
		if ext := filepath.Ext(name); ext != ".yaml" && ext != ".yml" {
			continue
		}

		b, err := fs.ReadFile(fsys, filepath.Join(dir, name))
		if err != nil {
			return nil, err
		}

		var ex Exercise
		if err := yaml.Unmarshal(b, &ex); err != nil {
			return nil, fmt.Errorf("%s: %w", name, err)
		}

		if ex.Slug == "" {
			return nil, fmt.Errorf("%s: missing slug", name)
		}

		out = append(out, ex)
	}

	return out, nil
}

// Fallback catalog in case of errors
// or no embedded exercises found.
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

// Discover local exercises from "exercises" directory
// Returns a slice of Exercise structs.
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

// List all exercises from catalog or local exercises.
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

// Get exercise by slug from catalog or local exercises.
func Get(slug string) (Exercise, error) {
	// Search catalog
	for _, ex := range catalog().Concepts {
		if ex.Slug == slug {
			return ex, nil
		}
	}
	for _, ex := range catalog().Projects {
		if ex.Slug == slug {
			return ex, nil
		}
	}

	// Search local
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

// Reset exercise template for a given slug
// from embedded FS to local exercises dir.
func Reset(ex Exercise) error {
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

// Initialize all exercises from embedded templates.
func InitAll() error {
	for _, ex := range catalog().Concepts {
		if err := copyExerciseTemplate(ex.Slug); err != nil {
			return err
		}
	}
	for _, ex := range catalog().Projects {
		if err := copyExerciseTemplate(ex.Slug); err != nil {
			return err
		}
	}
	return nil
}

// Copy exercise template from embedded FS to local exercises dir.
func copyExerciseTemplate(slug string) error {
	targetDir := filepath.Join("exercises", slug)

	// Clean slate
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
