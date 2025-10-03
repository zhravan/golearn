package taskscheduler

import "time"

// RunAfter runs a task once after the given delay.
func RunAfter(delay time.Duration, task func()) {
	time.Sleep(delay)
	task()
}

// RunEvery runs a task multiple times with the given delay between each run.
func RunEvery(delay time.Duration, count int, task func()) {
	for i := 0; i < count; i++ {
		time.Sleep(delay)
		task()
	}
}
