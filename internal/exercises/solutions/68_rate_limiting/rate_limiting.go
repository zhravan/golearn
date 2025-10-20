package rate_limiting

import (
	"sync"
	"time"
)

// RateLimiter limits the number of requests allowed per key within a given time interval.
type RateLimiter struct {
	mu         sync.Mutex
	limit      int
	interval   time.Duration
	timestamps map[string][]time.Time
}

// NewRateLimiter returns a new RateLimiter with the given limit and interval.
func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:      limit,
		interval:   interval,
		timestamps: make(map[string][]time.Time),
	}
}

// Allow checks if a request for the given key is allowed.
// It removes expired timestamps and enforces the request limit.
func (r *RateLimiter) Allow(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-r.interval)

	// Filter out timestamps that are outside the allowed interval
	validTimestamps := make([]time.Time, 0, len(r.timestamps[key]))
	for _, t := range r.timestamps[key] {
		if t.After(windowStart) {
			validTimestamps = append(validTimestamps, t)
		}
	}

	// Update the timestamps with only valid ones
	r.timestamps[key] = validTimestamps

	// Check if the current request is allowed
	if len(validTimestamps) < r.limit {
		// Add the current timestamp and allow the request
		r.timestamps[key] = append(r.timestamps[key], now)
		return true
	}

	// Otherwise, deny the request
	return false
}

// Reset clears the request history for a given key.
func (r *RateLimiter) Reset(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.timestamps, key)
}