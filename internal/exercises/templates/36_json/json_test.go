package json

import "testing"

func TestMarshalPerson(t *testing.T) {
	p := Person{Name: "golearn", Email: "golearn@example.com"}

	jsonStr, err := MarshalPerson(p)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := `{"name":"golearn","email":"golearn@example.com"}`
	if jsonStr != expected {
		t.Errorf("Expected %s, got %s", expected, jsonStr)
	}
}

func TestUnmarshalPerson(t *testing.T) {
	jsonStr := `{"name":"golearn","email":"golearn@example.com"}`

	p, err := UnmarshalPerson(jsonStr)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if p.Name != "golearn" {
		t.Errorf("Expected name 'golearn', got %s", p.Name)
	}
	if p.Email != "golearn@example.com" {
		t.Errorf("Expected email 'golearn@example.com', got %s", p.Email)
	}
}