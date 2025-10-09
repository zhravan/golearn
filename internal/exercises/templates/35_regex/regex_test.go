package regex

import "testing"

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"user@domain.com", true},
		{"test.email@example.org", true},
		{"invalid-email", false},
		{"@domain.com", false},
		{"user@", false},
		{"user@domain", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsValidEmail(test.email)
		if result != test.expected {
			t.Errorf("IsValidEmail(%q) = %v, want %v", test.email, result, test.expected)
		}
	}
}

func TestExtractNumbers(t *testing.T) {
	tests := []struct {
		text     string
		expected []string
	}{
		{"I have 123 apples and 45.67 oranges", []string{"123", "45.67"}},
		{"The price is $99.99", []string{"99.99"}},
		{"No numbers here", []string{}},
		{"123", []string{"123"}},
		{"12.34.56", []string{"12.34", "56"}},
		{"", []string{}},
	}

	for _, test := range tests {
		result := ExtractNumbers(test.text)
		if len(result) != len(test.expected) {
			t.Errorf("ExtractNumbers(%q) length = %d, want %d", test.text, len(result), len(test.expected))
			continue
		}
		for i, num := range result {
			if num != test.expected[i] {
				t.Errorf("ExtractNumbers(%q)[%d] = %q, want %q", test.text, i, num, test.expected[i])
			}
		}
	}
}

func TestReplaceVowels(t *testing.T) {
	tests := []struct {
		text     string
		expected string
	}{
		{"hello", "h*ll*"},
		{"HELLO", "H*LL*"},
		{"Hello World", "H*ll* W*rld"},
		{"bcdfg", "bcdfg"},
		{"", ""},
		{"aeiou", "*****"},
		{"AEIOU", "*****"},
	}

	for _, test := range tests {
		result := ReplaceVowels(test.text)
		if result != test.expected {
			t.Errorf("ReplaceVowels(%q) = %q, want %q", test.text, result, test.expected)
		}
	}
}

func TestIsPhoneNumber(t *testing.T) {
	tests := []struct {
		phone    string
		expected bool
	}{
		{"(123) 456-7890", true},
		{"(555) 123-4567", true},
		{"(000) 000-0000", true},
		{"123-456-7890", false},
		{"(123)456-7890", false},
		{"(123) 4567890", false},
		{"123 456 7890", false},
		{"(123) 456-789", false},
		{"(12) 456-7890", false},
		{"", false},
	}

	for _, test := range tests {
		result := IsPhoneNumber(test.phone)
		if result != test.expected {
			t.Errorf("IsPhoneNumber(%q) = %v, want %v", test.phone, result, test.expected)
		}
	}
}
