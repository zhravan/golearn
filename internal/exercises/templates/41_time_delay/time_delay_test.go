package timedelay

import (
	"testing"
	"time"
)

func TestWaitFor(t *testing.T) {
	start := time.Now()
	WaitFor(100) // wait 100 milliseconds
	elapsed := time.Since(start).Milliseconds()

	if elapsed < 90 {
		t.Fatalf("WaitFor(100) waited too short: %dms", elapsed)
	}
}

func TestNotifyAfter(t *testing.T) {
	start := time.Now()
	ch := NotifyAfter(100)

	select {
	case <-ch:
		elapsed := time.Since(start).Milliseconds()
		if elapsed < 90 {
			t.Fatalf("NotifyAfter(100) sent too early: %dms", elapsed)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("NotifyAfter(100) did not send on time")
	}
}
