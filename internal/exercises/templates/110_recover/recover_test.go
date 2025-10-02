package recover_exercise

import (
	"testing"
)

func TestRun_NoPanic(t *testing.T) {
	t.Parallel()
	result := Run(10)
	if result != nil {
		t.Errorf("Run(10) should not panic and should return nil. Got: %v", result)
	}
}

func TestRun_WithPanic(t *testing.T) {
	t.Parallel()
	expectedPanicValue := "input cannot be negative"
	result := Run(-5)

	if result == nil {
		t.Errorf("Run(-5) should panic and recover, returning the panic value. Got nil")
	}

	// Check if the recovered value is what we expected
	if resultAsString, ok := result.(string); !ok || resultAsString != expectedPanicValue {
		t.Errorf("Run(-5) recovered with unexpected value. Expected: %q, Got: %v (%T)", expectedPanicValue, result, result)
	}
}
