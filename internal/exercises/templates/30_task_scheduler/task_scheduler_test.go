package taskscheduler

import (
	"testing"
	"time"
)

// TestRunAfter checks if the task runs exactly once after a short delay.
func TestRunAfter(t *testing.T) {
	counter := 0

	RunAfter(10*time.Millisecond, func() {
		counter++
	})

	if counter != 1 {
		t.Errorf("expected task to run once, got %d", counter)
	}
}

// TestRunEvery checks if the task runs the correct number of times.
func TestRunEvery(t *testing.T) {
	counter := 0

	RunEvery(5*time.Millisecond, 3, func() {
		counter++
	})

	if counter != 3 {
		t.Errorf("expected task to run 3 times, got %d", counter)
	}
}
