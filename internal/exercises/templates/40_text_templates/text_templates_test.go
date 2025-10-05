package texttemplates

import (
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	tmpl := "Name: {{.Name}}, Age: {{.Age}}"
	data := map[string]interface{}{
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
}

func TestFormatUserGreeting(t *testing.T) {
	result, err := FormatUserGreeting("Bob")
	if err != nil {
		t.Errorf("FormatUserGreeting returned error: %v", err)
	}
	expected := "Hello, Bob!"
	if result != expected {
		t.Errorf("FormatUserGreeting = %q, want %q", result, expected)
	}
}
