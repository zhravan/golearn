package basic_key_value_store

import (
	"os"
	"testing"
)

func TestNewKeyValueStore(t *testing.T) {
	filepath := "test_kv_store.txt"
	s := NewKeyValueStore(filepath)
	if s == nil {
		t.Errorf("Expected a new KeyValueStore, got nil")
	}
	if s.filepath != filepath {
		t.Errorf("Expected filepath %s, got %s", filepath, s.filepath)
	}
	if s.data == nil {
		t.Errorf("Expected data map to be initialized, got nil")
	}
}

func TestSetAndGet(t *testing.T) {
	filepath := "test_set_get.txt"
	defer os.Remove(filepath)

	s := NewKeyValueStore(filepath)
	s.Set("name", "Alice")
	s.Set("age", "30")

	val, err := s.Get("name")
	if err != nil || val != "Alice" {
		t.Errorf("Expected 'Alice', got %q, error: %v", val, err)
	}

	val, err = s.Get("age")
	if err != nil || val != "30" {
		t.Errorf("Expected '30', got %q, error: %v", val, err)
	}

	_, err = s.Get("nonexistent")
	if err == nil {
		t.Errorf("Expected error for non-existent key, got nil")
	}
}

func TestDelete(t *testing.T) {
	filepath := "test_delete.txt"
	defer os.Remove(filepath)

	s := NewKeyValueStore(filepath)
	s.Set("key1", "value1")

	err := s.Delete("key1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = s.Get("key1")
	if err == nil {
		t.Errorf("Expected error for deleted key, got nil")
	}

	err = s.Delete("nonexistent")
	if err == nil {
		t.Errorf("Expected error for deleting non-existent key, got nil")
	}
}

func TestLoadAndSave(t *testing.T) {
	filepath := "test_load_save.txt"
	defer os.Remove(filepath)

	// Create a store and save it
	s1 := NewKeyValueStore(filepath)
	s1.Set("city", "New York")
	s1.Set("country", "USA")
	err := s1.Save()
	if err != nil {
		t.Fatalf("Failed to save store: %v", err)
	}

	// Load into a new store
	s2 := NewKeyValueStore(filepath)
	err = s2.Load()
	if err != nil {
		t.Fatalf("Failed to load store: %v", err)
	}

	val, err := s2.Get("city")
	if err != nil || val != "New York" {
		t.Errorf("Expected 'New York', got %q, error: %v", val, err)
	}

	val, err = s2.Get("country")
	if err != nil || val != "USA" {
		t.Errorf("Expected 'USA', got %q, error: %v", val, err)
	}
}

