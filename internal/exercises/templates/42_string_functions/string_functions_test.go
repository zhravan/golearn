package string_functions

import "testing"

func TestContains(t *testing.T) {
	tests := []struct {
		s      string
		substr string
		want   bool
	}{
		{"hello world", "world", true},
		{"hello world", "hello", true},
		{"hello world", "foo", false},
		{"", "", true},
		{"hello", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.s+" contains "+tt.substr, func(t *testing.T) {
			if got := Contains(tt.s, tt.substr); got != tt.want {
				t.Errorf("Contains(%q, %q) = %v, want %v", tt.s, tt.substr, got, tt.want)
			}
		})
	}
}

func TestHasPrefix(t *testing.T) {
	tests := []struct {
		s      string
		prefix string
		want   bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", false},
		{"", "", true},
		{"hello", "hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.s+" has prefix "+tt.prefix, func(t *testing.T) {
			if got := HasPrefix(tt.s, tt.prefix); got != tt.want {
				t.Errorf("HasPrefix(%q, %q) = %v, want %v", tt.s, tt.prefix, got, tt.want)
			}
		})
	}
}

func TestHasSuffix(t *testing.T) {
	tests := []struct {
		s      string
		suffix string
		want   bool
	}{
		{"hello world", "world", true},
		{"hello world", "hello", false},
		{"", "", true},
		{"hello", "hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.s+" has suffix "+tt.suffix, func(t *testing.T) {
			if got := HasSuffix(tt.s, tt.suffix); got != tt.want {
				t.Errorf("HasSuffix(%q, %q) = %v, want %v", tt.s, tt.suffix, got, tt.want)
			}
		})
	}
}

func TestIndex(t *testing.T) {
	tests := []struct {
		s      string
		substr string
		want   int
	}{
		{"hello world", "world", 6},
		{"hello world", "hello", 0},
		{"hello world", "foo", -1},
		{"", "", 0},
		{"hello", "l", 2},
	}

	for _, tt := range tests {
		t.Run("index of "+tt.substr+" in "+tt.s, func(t *testing.T) {
			if got := Index(tt.s, tt.substr); got != tt.want {
				t.Errorf("Index(%q, %q) = %v, want %v", tt.s, tt.substr, got, tt.want)
			}
		})
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"hello", "HELLO"},
		{"Hello World", "HELLO WORLD"},
		{"", ""},
		{"123", "123"},
		{"héllo", "HÉLLO"},
	}

	for _, tt := range tests {
		t.Run("ToUpper("+tt.s+")", func(t *testing.T) {
			if got := ToUpper(tt.s); got != tt.want {
				t.Errorf("ToUpper(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestToLower(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"HELLO", "hello"},
		{"Hello World", "hello world"},
		{"", ""},
		{"123", "123"},
		{"HÉLLO", "héllo"},
	}

	for _, tt := range tests {
		t.Run("ToLower("+tt.s+")", func(t *testing.T) {
			if got := ToLower(tt.s); got != tt.want {
				t.Errorf("ToLower(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}

func TestTrimSpace(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"  hello  ", "hello"},
		{"  hello world  ", "hello world"},
		{"", ""},
		{"hello", "hello"},
		{"   ", ""},
		{"\n\thello\t\n", "hello"},
	}

	for _, tt := range tests {
		t.Run("TrimSpace("+tt.s+")", func(t *testing.T) {
			if got := TrimSpace(tt.s); got != tt.want {
				t.Errorf("TrimSpace(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
