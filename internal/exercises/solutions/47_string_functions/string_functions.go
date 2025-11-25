package string_functions

import "strings"

// Contains checks if substr is within s
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// HasPrefix tests whether the string s begins with prefix
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix tests whether the string s ends with suffix
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s
func Index(s, substr string) int {
	return strings.Index(s, substr)
}

// ToUpper returns s with all Unicode letters mapped to their upper case
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower returns s with all Unicode letters mapped to their lower case
func ToLower(s string) string {
	return strings.ToLower(s)
}

// TrimSpace returns s with all leading and trailing white space removed
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}
