package timedelay

// Template with TODOs for teammates to implement using time.Now and time.Add.

import "time"

// WaitFor pauses execution until roughly ms milliseconds have passed.
func WaitFor(ms int) {
	// TODO: compute a target time using time.Now().Add and wait until reached
}

// NotifyAfter waits for roughly ms milliseconds then sends true on a channel.
func NotifyAfter(ms int) chan bool {
	// TODO: start a goroutine that checks time.Now against a target and sends on the channel
	return nil
}

// WaitUntil blocks until the provided target time is reached (or returns if in the past).
func WaitUntil(target time.Time) {
	// TODO: loop until time.Now() is not before target
}

// NotifyAt sends true on a channel when the provided target time is reached.
func NotifyAt(target time.Time) chan bool {
	// TODO: return a channel that will receive true when the target time arrives
	return nil
}

// ElapsedMillis returns how many milliseconds have passed since t.
func ElapsedMillis(t time.Time) int64 {
	// TODO: use time.Since to compute milliseconds
	return 0
}

// WaitForOrTimeout waits for ms milliseconds but returns an error if waiting exceeds timeoutMs.
func WaitForOrTimeout(ms int, timeoutMs int) error {
	// TODO: compute target and deadline with time.Now().Add and return error on timeout
	return nil
}

// ScheduleAfter runs fn after roughly ms milliseconds and returns a channel closed when fn finishes.
func ScheduleAfter(ms int, fn func()) chan struct{} {
	// TODO: start a goroutine that waits until target, runs fn, then closes the done channel
	return nil
}
