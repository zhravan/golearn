package structs

import "testing"

func TestNewPerson(t *testing.T) {
	p := NewPerson("Alice", 30)
	if p.Name != "Alice" {
		t.Errorf("Expected Name Alice, got %s", p.Name)
	}
	if p.Age != 30 {
		t.Errorf("Expected Age 30, got %d", p.Age)
	}
}
