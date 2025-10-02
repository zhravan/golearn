package string_formatting

import (
	"testing"
)

func TestStringFormats(t *testing.T) {
	t.Run("Name should be formatted as a string!", func(t *testing.T) {
		got := FormatName()
		want := "Name: \"John\""

		if got != want {
			t.Errorf("Expected %q, got %q", want, got)
		}
	})

	t.Run("Age should formatted as a digit!", func(t *testing.T) {
		got := FormatAge()
		want := "Age: 17"

		if got != want {
			t.Errorf("Expected %q got %q", want, got)
		}
	})

	t.Run("GPA should be formatted for floating point number with 2 decimal places!", func(t *testing.T) {
		got := FormatGpa()
		want := "GPA: 3.75"

		if got != want {
			t.Errorf("Expected %q got %q", want, got)
		}
	})
}
