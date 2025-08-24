package text_analyzer

import (
	"strings"
	"unicode/utf8"
)

func CountCharacters(text string) int {
	return utf8.RuneCountInString(text)
}

func CountWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func CountUniqueWords(text string) int {
	words := strings.Fields(text)
	uniqueWords := make(map[string]bool)
	for _, word := range words {
		uniqueWords[strings.ToLower(word)] = true
	}
	return len(uniqueWords)
}
