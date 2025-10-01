package basic_key_value_store

import (
	"errors"
	"path/filepath"
	"testing"
)

func TestLoadAndSave(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "test_store.db")

	store := NewKeyValueStore(filename)
	store.Set("city", "New York")
	store.Set("country", "USA")

	if err := store.Save(); err != nil {
		t.Fatalf("unexpected save error: %v", err)
	}

	// load into a new store
	other := NewKeyValueStore(filename)
	if err := other.Load(); err != nil {
		t.Fatalf("unexpected load error: %v", err)
	}

	if v, err := other.Get("city"); v != "New York" || err != nil {
		t.Errorf("expected New York, got %q, err=%v", v, err)
	}
	if v, err := other.Get("country"); v != "USA" || err != nil {
		t.Errorf("expected USA, got %q, err=%v", v, err)
	}
}

func TestMissingKey(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "test_store_missing.db")

	store := NewKeyValueStore(filename)
	store.Set("name", "Alice")

	_, err := store.Get("notthere")
	if !errors.Is(err, ErrKeyNotFound) {
		t.Errorf("expected ErrKeyNotFound, got %v", err)
	}
}

func TestDelete(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	filename := filepath.Join(dir, "test_store.db")

	store := NewKeyValueStore(filename)
	store.Set("k", "v")

	if err := store.Delete("k"); err != nil {
		t.Fatalf("unexpected delete error: %v", err)
	}
	if _, err := store.Get("k"); !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("expected ErrKeyNotFound after delete, got %v", err)
	}
	if err := store.Delete("k"); !errors.Is(err, ErrKeyNotFound) {
		t.Fatalf("expected ErrKeyNotFound on second delete, got %v", err)
	}
}

func TestLoadMissingFileIsNoop(t *testing.T) {
	t.Parallel()
	dir := t.TempDir()
	filename := filepath.Join(dir, "does_not_exist.db")

	store := NewKeyValueStore(filename)
	if err := store.Load(); err != nil {
		t.Fatalf("expected nil error on missing file, got %v", err)
	}
}
