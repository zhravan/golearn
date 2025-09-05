package basic_key_value_store

import (
	"os"
	"testing"
)

// helper to clean up after tests
func cleanup(filename string) {
	_ = os.Remove(filename)
}

func TestLoadAndSave(t *testing.T) {
	filename := "test_store.db"
	defer cleanup(filename)

	store := NewKeyValueStore(filename)

	// Save data
	store.Set("city", "New York")
	store.Set("country", "USA")

	// Persist to file
	if err := store.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}

	// Create a new store instance to load from file
	store2 := NewKeyValueStore(filename)

	// Reload data
	if err := store2.Load(); err != nil {
		t.Fatalf("failed to load: %v", err)
	}

	// Verify city
	val, err := store2.Get("city")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != "New York" {
		t.Errorf("Expected %q, got %q", "New York", val)
	}

	// Verify country
	val, err = store2.Get("country")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if val != "USA" {
		t.Errorf("Expected %q, got %q", "USA", val)
	}
}

func TestMissingKey(t *testing.T) {
	filename := "test_store_missing.db"
	defer cleanup(filename)

	store := NewKeyValueStore(filename)

	_, err := store.Get("notthere")
	if err == nil {
		t.Errorf("expected error for missing key, got nil")
	}
}
