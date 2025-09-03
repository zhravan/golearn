package basic_key_value_store

import (
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
    return &KeyValueStore{}
}

func (s *KeyValueStore) Load() error {
    // TODO: load key/value pairs from file if present
    return nil
}

func (s *KeyValueStore) Save() error {
    // TODO: write pairs to file
    return nil
}

func (s *KeyValueStore) Set(key, value string) {
    // TODO: set key to value
}

func (s *KeyValueStore) Get(key string) (string, *StoreError) {
    // TODO: get value or return error when missing
    return "", nil
}

func (s *KeyValueStore) Delete(key string) *StoreError {
    // TODO: delete key or return error when missing
    return nil
}
