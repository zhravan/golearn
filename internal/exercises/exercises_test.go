package exercises

import (
	"strings"
	"testing"
	"testing/fstest"

	"gopkg.in/yaml.v3"
)

// Helper to build YAML quickly
func y(v any) []byte {
	b, _ := yaml.Marshal(v)
	return b
}

func TestLoadExercisesDir_LoadsValidYAML(t *testing.T) {
	fsys := fstest.MapFS{
		"Catalog/Concepts/a.yaml": &fstest.MapFile{
			Data: y(Exercise{
				Slug:      "02_vars",
				Title:     "Variables",
				TestRegex: ".*",
				Hints:     []string{"x := 1"},
			}),
		},
		"Catalog/Concepts/b.yaml": &fstest.MapFile{
			Data: y(Exercise{
				Slug:      "01_hello",
				Title:     "Hello",
				TestRegex: ".*",
			}),
		},
	}

	items, err := loadExercisesDir(fsys, "Catalog/Concepts")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(items) != 2 {
		t.Fatalf("expected 2 items, got %d", len(items))
	}

	// Unsorted read → sorted in loadCatalogFromFS, but loadExercisesDir returns them sorted by slug.
	if items[0].Slug != "01_hello" || items[1].Slug != "02_vars" {
		t.Fatalf("items not sorted properly: %v", items)
	}
}

func TestLoadExercisesDir_RejectsEmptySlug(t *testing.T) {
	fsys := fstest.MapFS{
		"Catalog/Concepts/bad.yaml": &fstest.MapFile{
			Data: []byte("title: NoSlug"),
		},
	}

	_, err := loadExercisesDir(fsys, "Catalog/Concepts")
	if err == nil || !strings.Contains(err.Error(), "missing slug") {
		t.Fatalf("expected missing slug error, got: %v", err)
	}
}

func TestLoadExercisesDir_IgnoresNonYAML(t *testing.T) {
	fsys := fstest.MapFS{
		"Catalog/Concepts/readme.txt": &fstest.MapFile{Data: []byte("ignore me")},
		"Catalog/Concepts/x.yaml":     &fstest.MapFile{Data: y(Exercise{Slug: "01_x"})},
	}

	items, err := loadExercisesDir(fsys, "Catalog/Concepts")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 1 {
		t.Fatalf("expected 1 YAML item, got %d", len(items))
	}
}

func TestLoadCatalogFromFS_LoadsConceptsAndProjects(t *testing.T) {
	fsys := fstest.MapFS{
		"Catalog/Concepts/a.yaml": &fstest.MapFile{
			Data: y(Exercise{Slug: "02_b"}),
		},
		"Catalog/Concepts/b.yaml": &fstest.MapFile{
			Data: y(Exercise{Slug: "01_a"}),
		},
		"Catalog/Projects/c.yaml": &fstest.MapFile{
			Data: y(Exercise{Slug: "10_p"}),
		},
	}

	cat, err := loadCatalogFromFS(fsys)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(cat.Concepts) != 2 {
		t.Fatalf("expected 2 concepts, got %d", len(cat.Concepts))
	}
	if len(cat.Projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(cat.Projects))
	}

	// Check sorted order
	if cat.Concepts[0].Slug != "01_a" || cat.Concepts[1].Slug != "02_b" {
		t.Fatalf("concepts not sorted: %+v", cat.Concepts)
	}
}

func TestLoadCatalogFromFS_MissingDirsAreIgnored(t *testing.T) {
	fsys := fstest.MapFS{
		"Catalog/Concepts/x.yaml": &fstest.MapFile{
			Data: y(Exercise{Slug: "01_a"}),
		},
	}

	cat, err := loadCatalogFromFS(fsys)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(cat.Projects) != 0 {
		t.Fatalf("expected zero projects")
	}
	if len(cat.Concepts) != 1 {
		t.Fatalf("expected one concept")
	}
}

func TestLoadCatalogFromFS_InvalidYAML(t *testing.T) {
	fsys := fstest.MapFS{
		"Catalog/Concepts/bad.yaml": &fstest.MapFile{
			Data: []byte("{invalid yaml} : ???"),
		},
	}

	_, err := loadCatalogFromFS(fsys)
	if err == nil {
		t.Fatal("expected YAML error, got nil")
	}
}

func TestFallbackCatalog(t *testing.T) {
	f := fallbackCatalog()
	if len(f.Concepts) != 1 {
		t.Fatalf("fallback should have exactly 1 concept")
	}
	if f.Concepts[0].Slug != "01_hello" {
		t.Fatalf("fallback wrong slug: %s", f.Concepts[0].Slug)
	}
}

func TestCatalog_UsesDirectoryLoader(t *testing.T) {
	// We verify that the singleton `catalog()` function actually calls our custom loader logic.
	fsys := fstest.MapFS{
		"Catalog/Concepts/a.yaml": &fstest.MapFile{
			Data: y(Exercise{Slug: "01_test"}),
		},
	}

	// Use the exported helper to swap the loader with one that uses our test filesystem
	WithTestCatalogLoader(func() (Catalog, error) {
		return loadCatalogFromFS(fsys)
	}, func() {
		cat := catalog()
		if len(cat.Concepts) != 1 || cat.Concepts[0].Slug != "01_test" {
			t.Fatalf("catalog() did not load from the injected FS")
		}
	})
}

func TestDiscoverLocal_NoDir(t *testing.T) {
	_, err := discoverLocal()
	if err != nil {
		t.Fatalf("should not error if exercises directory is missing")
	}
}

func TestCatalogOverride(t *testing.T) {
	WithTestCatalogLoader(func() (Catalog, error) {
		return Catalog{
			Concepts: []Exercise{{Slug: "01_mock"}},
		}, nil
	}, func() {
		c := catalog()

		if len(c.Concepts) != 1 || c.Concepts[0].Slug != "01_mock" {
			t.Fatalf("expected mock catalog, got: %+v", c)
		}
	})
}
