package regex

// Task:
// Implement regular expression functions using Go's regexp package.
//
// 1. Implement IsValidEmail to validate email addresses using regex.
// 2. Implement ExtractNumbers to extract all numbers from a string.
// 3. Implement ReplaceVowels to replace all vowels with asterisks.
// 4. Implement IsPhoneNumber to validate phone numbers in format (XXX) XXX-XXXX.

// IsValidEmail should return true if the email address is valid.
// A valid email should have the format: user@domain.com
func IsValidEmail(email string) bool {
	// TODO: implement using regexp.MustCompile and MatchString
	return false
}

// ExtractNumbers should return all numbers found in the input string.
// Numbers can be integers or decimals (e.g., "123", "45.67").
func ExtractNumbers(text string) []string {
	// TODO: implement using regexp.FindAllString
	return nil
}

// ReplaceVowels should replace all vowels (a, e, i, o, u) with asterisks.
// Case-insensitive replacement.
func ReplaceVowels(text string) string {
	// TODO: implement using regexp.MustCompile and ReplaceAllString
	return ""
}

// IsPhoneNumber should return true if the phone number matches format (XXX) XXX-XXXX.
// Example: (123) 456-7890
func IsPhoneNumber(phone string) bool {
	// TODO: implement using regexp.MustCompile and MatchString
	return false
}
