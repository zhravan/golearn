package timedelay

// Real implementations using time.Now, time.Add and comparisons for the exercise.

import (
	"errors"
	"time"
)

// WaitFor pauses execution until roughly ms milliseconds have passed.
func WaitFor(ms int) {
	target := time.Now().Add(time.Duration(ms) * time.Millisecond)
	for time.Now().Before(target) {
		time.Sleep(1 * time.Millisecond)
	}
}

// NotifyAfter returns a channel that receives true after roughly ms milliseconds.
func NotifyAfter(ms int) chan bool {
	ch := make(chan bool, 1)
	target := time.Now().Add(time.Duration(ms) * time.Millisecond)

	go func() {
		for time.Now().Before(target) {
			time.Sleep(1 * time.Millisecond)
		}
		ch <- true
	}()

	return ch
}

// WaitUntil blocks until the provided target time is reached (returns immediately if in the past).
func WaitUntil(target time.Time) {
	for time.Now().Before(target) {
		time.Sleep(1 * time.Millisecond)
	}
}

// NotifyAt returns a channel that receives true when target time is reached (sends immediately if in the past).
func NotifyAt(target time.Time) chan bool {
	ch := make(chan bool, 1)

	go func() {
		for time.Now().Before(target) {
			time.Sleep(1 * time.Millisecond)
		}
		ch <- true
	}()

	return ch
}

// ElapsedMillis returns how many milliseconds have elapsed since t.
func ElapsedMillis(t time.Time) int64 {
	return time.Since(t).Milliseconds()
}

// WaitForOrTimeout waits up to ms milliseconds but fails if total waiting exceeds timeoutMs.
func WaitForOrTimeout(ms int, timeoutMs int) error {
	target := time.Now().Add(time.Duration(ms) * time.Millisecond)
	deadline := time.Now().Add(time.Duration(timeoutMs) * time.Millisecond)

	for {
		now := time.Now()
		if !now.Before(target) {
			return nil
		}
		if !now.Before(deadline) {
			return errors.New("timeout exceeded before target reached")
		}
		time.Sleep(1 * time.Millisecond)
	}
}

// ScheduleAfter runs fn after roughly ms milliseconds and returns a channel closed when fn finishes.
func ScheduleAfter(ms int, fn func()) chan struct{} {
	done := make(chan struct{})
	target := time.Now().Add(time.Duration(ms) * time.Millisecond)

	go func() {
		for time.Now().Before(target) {
			time.Sleep(1 * time.Millisecond)
		}
		fn()
		close(done)
	}()

	return done
}
