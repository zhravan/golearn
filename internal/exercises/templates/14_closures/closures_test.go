package closures

import "testing"

func TestClosure(t *testing.T) {
	next := closure()
	if next() != 1 {
		t.Errorf("Expected 1, got %d", next())
	}
	if next() != 2 {
		t.Errorf("Expected 2, got %d", next())
	}

	anotherNext := closure()
	if anotherNext() != 1 {
		t.Errorf("Expected 1, got %d", anotherNext())
	}
}
