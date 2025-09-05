package basic_key_value_store

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

// TODO:
// - Implement a basic persistent key-value store:
//   - In-memory map protected with RWMutex.
//   - Load: read key=value pairs from a file if present.
//   - Save: write all pairs to a file.
//   - Set/Get/Delete operations; Get/Delete error on missing keys.

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

func NewKeyValueStore(filepath string) *KeyValueStore {
	// TODO: initialize the store
	return &KeyValueStore{
		data:     make(map[string]string),
		filepath: filepath,
	}
}

func (s *KeyValueStore) Load() error {
	// TODO: load key/value pairs from file if present
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open(s.filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			s.data[parts[0]] = parts[1]
		}
	}

	return scanner.Err()
}

func (s *KeyValueStore) Save() error {
	// TODO: write pairs to file
	s.mu.Lock()
	defer s.mu.Unlock()
	file, err := os.Create(s.filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for k, v := range s.data {
		_, err := fmt.Fprintf(writer, "%s=%s\n", k, v)
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

func (s *KeyValueStore) Set(key, value string) {
	// TODO: set key to value
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *KeyValueStore) Get(key string) (string, *StoreError) {
	// TODO: get value or return error when missing
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	if !ok {
		return "", &StoreError{Message: "key not found"}
	}
	return val, nil
}

func (s *KeyValueStore) Delete(key string) *StoreError {
	// TODO: delete key or return error when missing
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[key]; !ok {
		return &StoreError{Message: "key not found"}
	}
	delete(s.data, key)
	return nil
}
