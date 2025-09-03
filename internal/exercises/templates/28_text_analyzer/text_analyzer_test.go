package text_analyzer

import "testing"

func TestCountCharacters(t *testing.T) {
	text := "Hello, 世界!"
	if CountCharacters(text) != 9 {
		t.Errorf("Expected 9 characters, got %d", CountCharacters(text))
	}
}

func TestCountWords(t *testing.T) {
	text := "Hello world, hello Go"
	if CountWords(text) != 4 {
		t.Errorf("Expected 4 words, got %d", CountWords(text))
	}
}

func TestCountUniqueWords(t *testing.T) {
	text := "Hello world, hello Go, world"
	if CountUniqueWords(text) != 3 {
		t.Errorf("Expected 3 unique words, got %d", CountUniqueWords(text))
	}
}
