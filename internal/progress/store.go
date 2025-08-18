package progress

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type state struct {
	Completed map[string]bool `json:"completed"`
}

func localStateFilePath() (string, bool) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", false
	}
	root := filepath.Join(cwd, ".golearn")
	if err := os.MkdirAll(root, 0o755); err != nil {
		return "", false
	}
	return filepath.Join(root, "progress.json"), true
}

func userStateFilePath() string {
	dir, _ := os.UserConfigDir()
	if dir == "" {
		dir = "."
	}
	root := filepath.Join(dir, "golearn")
	_ = os.MkdirAll(root, 0o755)
	return filepath.Join(root, "progress.json")
}

func stateFile() string {
	if p, ok := localStateFilePath(); ok {
		return p
	}
	return userStateFilePath()
}

func load() (state, error) {
	var s state
	b, err := os.ReadFile(stateFile())
	if errors.Is(err, os.ErrNotExist) {
		s.Completed = map[string]bool{}
		return s, nil
	}
	if err != nil {
		return s, err
	}
	if err := json.Unmarshal(b, &s); err != nil {
		return s, err
	}
	if s.Completed == nil {
		s.Completed = map[string]bool{}
	}
	return s, nil
}

func save(s state) error {
	b, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(stateFile(), b, 0o644)
}

func MarkCompleted(slug string) error {
	s, err := load()
	if err != nil {
		return err
	}
	s.Completed[slug] = true
	return save(s)
}

func IsCompleted(slug string) (bool, error) {
	s, err := load()
	if err != nil {
		return false, err
	}
	return s.Completed[slug], nil
}
