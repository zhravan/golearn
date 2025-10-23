package timers

import (
	"sync"
	"time"
)

// TimerManager manages timers that trigger callbacks after a specified duration.
//
// Each timer is identified by a string key and automatically runs its callback when it expires.
// Timers can be started, reset, or stopped safely from multiple goroutines.
//
// TODO:
// - Implement a TimerManager that stores timers in a map[string]*time.Timer.
// - Use a mutex to handle concurrent access safely.
// - Implement Start(key string, d time.Duration, fn func()) to start a timer that triggers fn after d.
// - If a timer already exists for a key, stop and replace it.
// - Implement Stop(key string) bool to stop a timer and remove it from the map.
// - Implement Reset(key string, d time.Duration) to reset a timer to fire again after d.
// - If Reset is called on a missing timer, do nothing.

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
	// TODO:
	// - Lock the mutex
	// - If a timer exists, stop and delete it
	// - Create a new timer using time.AfterFunc(d, wrapper)
	//   where wrapper: locks, deletes the timer from map, unlocks, then calls fn()
	// - Store it in the map
	// - Unlock the mutex
}

// Stop stops and removes the timer with the given key.
// It returns true if the timer existed and was stopped, false otherwise.
func (tm *TimerManager) Stop(key string) bool {
	// TODO:
	// - Lock the mutex
	// - Look up the timer
	// - If found, stop it, delete from map, unlock, return true
	// - Otherwise unlock and return false
	return false
}

// Reset resets the existing timer for the given key to expire after duration d again.
// If no timer exists for that key, do nothing.
func (tm *TimerManager) Reset(key string, d time.Duration) {
	// TODO:
	// - Lock the mutex
	// - Look up the timer
	// - If found, call timer.Reset(d)
	// - Unlock the mutex
}