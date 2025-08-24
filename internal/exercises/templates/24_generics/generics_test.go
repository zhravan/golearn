package generics

import "testing"

func TestIdentity(t *testing.T) {
	if identity(1) != 1 {
		t.Errorf("Expected 1, got %v", identity(1))
	}
	if identity("hello") != "hello" {
		t.Errorf("Expected hello, got %v", identity("hello"))
	}
}

func TestSumIntsOrFloats(t *testing.T) {
	if sumIntsOrFloats(1, 2) != 3 {
		t.Errorf("Expected 3, got %v", sumIntsOrFloats(1, 2))
	}
	if sumIntsOrFloats(1.0, 2.0) != 3.0 {
		t.Errorf("Expected 3.0, got %v", sumIntsOrFloats(1.0, 2.0))
	}
}
