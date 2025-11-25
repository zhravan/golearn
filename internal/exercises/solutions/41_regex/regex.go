package regex

import "regexp"

// IsValidEmail returns true if the email address is valid.
func IsValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(email)
}

// ExtractNumbers returns all numbers found in the input string.
func ExtractNumbers(text string) []string {
	pattern := `\d+\.?\d*`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllString(text, -1)
	if matches == nil {
		return []string{}
	}
	return matches
}

// ReplaceVowels replaces all vowels with asterisks.
func ReplaceVowels(text string) string {
	pattern := `[aeiouAEIOU]`
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, "*")
}

// IsPhoneNumber returns true if the phone number matches format (XXX) XXX-XXXX.
func IsPhoneNumber(phone string) bool {
	pattern := `^\(\d{3}\) \d{3}-\d{4}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(phone)
}
