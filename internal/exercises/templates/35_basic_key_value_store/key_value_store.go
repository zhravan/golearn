package basic_key_value_store

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

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
	return &KeyValueStore{
		data:     make(map[string]string),
		filepath: filepath,
	}
}

func (s *KeyValueStore) Load() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	file, err := os.Open(s.filepath)
	if os.IsNotExist(err) {
		return nil // No file, nothing to load
	}
	if err != nil {
		return &StoreError{Message: fmt.Sprintf("failed to open store file: %v", err)}
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
	return nil
}

func (s *KeyValueStore) Save() error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	file, err := os.Create(s.filepath)
	if err != nil {
		return &StoreError{Message: fmt.Sprintf("failed to create store file: %v", err)}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for k, v := range s.data {
		_, err := writer.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return &StoreError{Message: fmt.Sprintf("failed to write to store file: %v", err)}
		}
	}
	return writer.Flush()
}

func (s *KeyValueStore) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *KeyValueStore) Get(key string) (string, *StoreError) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	if !ok {
		return "", &StoreError{Message: fmt.Sprintf("key '%s' not found", key)}
	}
	return val, nil
}

func (s *KeyValueStore) Delete(key string) *StoreError {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.data[key]; !ok {
		return &StoreError{Message: fmt.Sprintf("key '%s' not found", key)}
	}
	delete(s.data, key)
	return nil
}

