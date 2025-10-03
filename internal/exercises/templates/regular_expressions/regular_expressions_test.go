package templates

import "testing"

func TestMatchRegex(t *testing.T) {
	tests := []struct {
		pattern string
		input   string
		want    bool
	}{
		{"^a.*z$", "abcz", true},
		{"^a.*z$", "abz", true},
		{"^a.*z$", "abc", false},
		{"[0-9]+", "123", true},
		{"[0-9]+", "abc", false},
		{"hello|world", "hello", true},
		{"hello|world", "world", true},
		{"hello|world", "hi", false},
	}
	for _, tt := range tests {
		got := MatchRegex(tt.pattern, tt.input)
		if got != tt.want {
			t.Errorf("MatchRegex(%q, %q) = %v; want %v", tt.pattern, tt.input, got, tt.want)
		}
	}
}
