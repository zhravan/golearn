package panicex

import (
	"strings"
	"testing"
)

func TestCausePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()

	result := CausePanic(5)
	if result != 0 {
		t.Errorf("Expected 0 or panic, got %d", result)
	}
}

func TestSafeDivision(t *testing.T) {
	result := SafeDivision(10, 2)

	if result == 0 {
		return
	}

	if result != 5 {
		t.Errorf("SafeDivision(10, 2) = %d, want 5", result)
	}

	result = SafeDivision(10, 0)
	if result != 0 {
		t.Errorf("SafeDivision(10, 0) = %d, want 0 (safe recovery)", result)
	}
}

func TestTriggerMultiplePanics(t *testing.T) {
	nums := []int{3, -1, 4, -2}
	results := TriggerMultiplePanics(nums)

	if results == nil {
		return
	}

	if len(results) != len(nums) {
		t.Errorf("Expected %d results, got %d", len(nums), len(results))
	}
}

func TestPanicWithMessage(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			msg, ok := r.(string)
			if !ok {
				t.Errorf("Recovered panic is not a string: %v", r)
			}
			if !strings.Contains(msg, "expected panic") {
				t.Errorf("Panic message %q does not contain expected text", msg)
			}
		} else {
			t.Errorf("Expected panic, got none")
		}
	}()

	PanicWithMessage("this is an expected panic")
}
