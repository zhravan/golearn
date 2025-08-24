package recursion

import "testing"

func TestFactorial(t *testing.T) {
	if factorial(0) != 1 {
		t.Errorf("Expected 1, got %d", factorial(0))
	}
	if factorial(1) != 1 {
		t.Errorf("Expected 1, got %d", factorial(1))
	}
	if factorial(5) != 120 {
		t.Errorf("Expected 120, got %d", factorial(5))
	}
}
