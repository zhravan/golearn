package basic_key_value_store

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

// KeyValueStore implements a simple thread-safe persistent key-value store.
type KeyValueStore struct {
	data     map[string]string
	filepath string
	mu       sync.RWMutex
}

type StoreError struct {
	Message string
}

func (e *StoreError) Error() string {
	return e.Message
}

// Sentinel error for missing keys.
var ErrKeyNotFound = &StoreError{Message: "key not found"}

// NewKeyValueStore initializes a new key-value store at the given filepath.
func NewKeyValueStore(filepath string) *KeyValueStore {
	return &KeyValueStore{
		data:     make(map[string]string),
		filepath: filepath,
	}
}

// Load replaces the current store with the contents of the file.
func (s *KeyValueStore) Load() error {
	file, err := os.Open(s.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0), 1024*1024) // allow long lines up to 1MB
	tmp := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			k := strings.TrimSpace(parts[0])
			v := strings.TrimSpace(parts[1])
			tmp[k] = v
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	s.mu.Lock()
	s.data = tmp
	s.mu.Unlock()
	return nil
}

// Save writes the store atomically to disk.
func (s *KeyValueStore) Save() error {
	// Snapshot under read lock
	s.mu.RLock()
	snapshot := make(map[string]string, len(s.data))
	for k, v := range s.data {
		snapshot[k] = v
	}
	s.mu.RUnlock()

	dir := filepath.Dir(s.filepath)
	tmp, err := os.CreateTemp(dir, ".kvtmp-*")
	if err != nil {
		return err
	}

	writer := bufio.NewWriter(tmp)
	for k, v := range snapshot {
		if _, err := fmt.Fprintf(writer, "%s=%s\n", k, v); err != nil {
			tmp.Close()
			_ = os.Remove(tmp.Name())
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		tmp.Close()
		_ = os.Remove(tmp.Name())
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		_ = os.Remove(tmp.Name())
		return err
	}
	if err := tmp.Close(); err != nil {
		_ = os.Remove(tmp.Name())
		return err
	}
	return os.Rename(tmp.Name(), s.filepath)
}

// Set updates or inserts a key.
func (s *KeyValueStore) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

// Get retrieves a value or ErrKeyNotFound.
func (s *KeyValueStore) Get(key string) (string, *StoreError) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return val, nil
}

// Delete removes a key or returns ErrKeyNotFound.
func (s *KeyValueStore) Delete(key string) *StoreError {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[key]; !ok {
		return ErrKeyNotFound
	}
	delete(s.data, key)
	return nil
}
