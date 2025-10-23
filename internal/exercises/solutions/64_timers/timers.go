package timers

import (
	"sync"
	"time"
)

type TimerManager struct {
	mu     sync.Mutex
	timers map[string]*time.Timer
}

// NewTimerManager returns a new TimerManager with an initialized timers map.
func NewTimerManager() *TimerManager {
	return &TimerManager{
		timers: make(map[string]*time.Timer),
	}
}

// Start starts a timer for the given key to run the provided callback after duration d.
// If a timer with the same key already exists, it should be stopped and replaced.
func (tm *TimerManager) Start(key string, d time.Duration, fn func()) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if oldTimer, ok := tm.timers[key]; ok {
		oldTimer.Stop()
		delete(tm.timers, key)
	}

	timer := time.AfterFunc(d, func() {
		tm.mu.Lock()
		delete(tm.timers, key)
		tm.mu.Unlock()

		fn()
	})

	tm.timers[key] = timer
}

// Stop stops and removes the timer with the given key.
// It returns true if the timer existed and was stopped, false otherwise.
func (tm *TimerManager) Stop(key string) bool {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	if timer, ok := tm.timers[key]; ok {
		timer.Stop()
		delete(tm.timers, key)
		return true
	}

	return false
}

// Reset resets the existing timer for the given key to expire after duration d again.
// If no timer exists for that key, do nothing.
func (tm *TimerManager) Reset(key string, d time.Duration) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	
	if timer, ok := tm.timers[key]; ok {
		timer.Reset(d)
	}
}