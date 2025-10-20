package timedelay

// Tests that validate the timing helpers with reasonable slack for CI.

import (
	"testing"
	"time"
)

func TestWaitFor(t *testing.T) {
	start := time.Now()
	WaitFor(100)
	el := time.Since(start).Milliseconds()
	if el < 90 {
		t.Fatalf("WaitFor(100) too short: %dms", el)
	}
	if el > 400 {
		t.Fatalf("WaitFor(100) too long: %dms", el)
	}
}

func TestNotifyAfter(t *testing.T) {
	start := time.Now()
	ch := NotifyAfter(100)

	select {
	case <-ch:
		el := time.Since(start).Milliseconds()
		if el < 90 {
			t.Fatalf("NotifyAfter(100) too early: %dms", el)
		}
		if el > 500 {
			t.Fatalf("NotifyAfter(100) too late: %dms", el)
		}
	case <-time.After(600 * time.Millisecond):
		t.Fatal("NotifyAfter(100) did not send")
	}
}

func TestWaitUntilAndNotifyAt(t *testing.T) {
	start := time.Now()
	target := time.Now().Add(120 * time.Millisecond)
	WaitUntil(target)
	el := time.Since(start).Milliseconds()
	if el < 100 {
		t.Fatalf("WaitUntil reached too early: %dms", el)
	}
	if el > 600 {
		t.Fatalf("WaitUntil took too long: %dms", el)
	}

	start2 := time.Now()
	target2 := time.Now().Add(100 * time.Millisecond)
	ch := NotifyAt(target2)
	select {
	case <-ch:
		el2 := time.Since(start2).Milliseconds()
		if el2 < 80 {
			t.Fatalf("NotifyAt too early: %dms", el2)
		}
		if el2 > 500 {
			t.Fatalf("NotifyAt too late: %dms", el2)
		}
	case <-time.After(600 * time.Millisecond):
		t.Fatal("NotifyAt did not send")
	}
}

func TestElapsedMillis(t *testing.T) {
	past := time.Now().Add(-150 * time.Millisecond)
	if ElapsedMillis(past) < 130 {
		t.Fatalf("ElapsedMillis too small: %dms", ElapsedMillis(past))
	}
}

func TestWaitForOrTimeout(t *testing.T) {
	if err := WaitForOrTimeout(100, 300); err != nil {
		t.Fatalf("expected nil, got %v", err)
	}
	if err := WaitForOrTimeout(300, 50); err == nil {
		t.Fatal("expected timeout error, got nil")
	}
}

func TestScheduleAfter(t *testing.T) {
	flag := make(chan bool, 1)
	fn := func() { flag <- true }

	complete := ScheduleAfter(120, fn)

	select {
	case <-flag:
	case <-time.After(700 * time.Millisecond):
		t.Fatal("scheduled function did not run")
	}

	select {
	case <-complete:
	case <-time.After(700 * time.Millisecond):
		t.Fatal("completion channel not closed")
	}
}
