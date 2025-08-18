package values

import "testing"

func TestFormatValues(t *testing.T) {
	got := FormatValues(3, 3.14, "go")
	want := "i=3 f=3.14 s=go"
	if got != want {
		t.Fatalf("FormatValues() = %q, want %q", got, want)
	}
}
