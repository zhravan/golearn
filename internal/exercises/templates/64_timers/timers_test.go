package timers

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestStartAndCallback(t *testing.T) {
	tm := NewTimerManager()
	var called atomic.Bool

	tm.Start("t1", 50*time.Millisecond, func() {
		called.Store(true)
	})

	time.Sleep(80 * time.Millisecond)

	if !called.Load() {
		t.Error("Expected callback to be called after timer duration")
	}
}

func TestStop(t *testing.T) {
	tm := NewTimerManager()
	var called atomic.Bool

	tm.Start("t2", 100*time.Millisecond, func() {
		called.Store(true)
	})

	stopped := tm.Stop("t2")
	if !stopped {
		t.Error("Expected Stop to return true for existing timer")
	}

	time.Sleep(150 * time.Millisecond)
	if called.Load() {
		t.Error("Expected callback not to be called after timer stopped")
	}
}

func TestReset(t *testing.T) {
	tm := NewTimerManager()
	var called atomic.Bool

	tm.Start("t3", 100*time.Millisecond, func() {
		called.Store(true)
	})

	time.Sleep(50 * time.Millisecond)
	tm.Reset("t3", 100*time.Millisecond)

	time.Sleep(70 * time.Millisecond)
	if called.Load() {
		t.Error("Expected callback not yet to be called after reset")
	}

	time.Sleep(50 * time.Millisecond)
	if !called.Load() {
		t.Error("Expected callback to be called after reset duration")
	}
}

func TestStopNonexistent(t *testing.T) {
	tm := NewTimerManager()
	if tm.Stop("nope") {
		t.Error("Expected Stop to return false for missing timer")
	}
}