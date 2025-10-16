package strings_runes

import "unicode/utf8"

// countRunes returns the number of runes (Unicode code points) in a string.
// Runes represent Unicode characters, which may be multi-byte.
// For example, "hello" has 5 runes, and "你好世界" has 4 runes (not 12 bytes).
func countRunes(s string) int {
	return utf8.RuneCountInString(s)
}

// reverseString reverses a string by rune, preserving multi-byte characters.
// It converts the string to a slice of runes, reverses the slice, then converts back.
// This ensures multi-byte Unicode characters like "你" are kept intact.
func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
