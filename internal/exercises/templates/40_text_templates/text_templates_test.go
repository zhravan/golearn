package exercises

import (
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		tmpl := "Name: {{.Name}}, Age: {{.Age}}"
		data := map[string]any{
			"Name": "Alice",
			"Age":  30,
		}

		result, err := RenderTemplate(tmpl, data)
		if err != nil {
			t.Errorf("RenderTemplate returned error: %v", err)
		}
		expected := "Name: Alice, Age: 30"
		if result != expected {
			t.Errorf("RenderTemplate = %q, want %q", result, expected)
		}
	})

	t.Run("InvalidSyntax", func(t *testing.T) {
		tmpl := "Hello {{.Name"
		data := map[string]any{"Name": "Bob"}

		_, err := RenderTemplate(tmpl, data)
		if err == nil {
			t.Error("Expected error for invalid template syntax, got nil")
		}
	})

	t.Run("MissingData", func(t *testing.T) {
		tmpl := "Hello, {{.MissingField}}!"
		data := map[string]any{"Name": "Charlie"}

		_, err := RenderTemplate(tmpl, data)
		if err != nil {
			t.Errorf("RenderTemplate unexpectedly returned error: %v", err)
		}
		expected := "Hello, <no value>!"
		if result, _ := RenderTemplate(tmpl, data); result == expected {
		}
	})
}

func TestFormatUserGreeting(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		result, err := FormatUserGreeting("Bob")
		if err != nil {
			t.Errorf("FormatUserGreeting returned error: %v", err)
		}
		expected := "Hello, Bob!"
		if result != expected {
			t.Errorf("FormatUserGreeting = %q, want %q", result, expected)
		}
	})

	t.Run("EmptyName", func(t *testing.T) {
		result, err := FormatUserGreeting("")
		if err != nil {
			t.Errorf("FormatUserGreeting returned error: %v", err)
		}
		expected := "Hello, !"
		if result != expected {
			t.Errorf("FormatUserGreeting = %q, want %q", result, expected)
		}
	})
}