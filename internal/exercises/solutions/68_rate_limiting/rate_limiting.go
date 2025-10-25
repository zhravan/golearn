package rate_limiting

import (
	"errors"
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
func NewRateLimiter(limit int, interval time.Duration) (*RateLimiter, error) {
	if limit <= 0 {
		return nil, errors.New("limit must be positive")
	}
	if interval <= 0 {
		return nil, errors.New("interval must be positive")
	}

	return &RateLimiter{
		limit:      limit,
		interval:   interval,
		timestamps: make(map[string][]time.Time),
	}, nil
}

// Allow checks if a request for the given key is allowed.
// It removes expired timestamps and enforces the request limit.
func (r *RateLimiter) Allow(key string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	windowStart := now.Add(-r.interval)

	// Filter out timestamps outside the interval
	validTimestamps := make([]time.Time, 0)
	for _, t := range r.timestamps[key] {
		if t.After(windowStart) {
			validTimestamps = append(validTimestamps, t)
		}
	}

	// Clean up keys if no valid timestamps remain
	if len(validTimestamps) == 0 {
		delete(r.timestamps, key)
	} else {
		r.timestamps[key] = validTimestamps
	}

	// Enforce limit
	if len(validTimestamps) < r.limit {
		r.timestamps[key] = append(r.timestamps[key], now)
		return true
	}

	return false
}

// Reset clears the request history for a given key.
func (r *RateLimiter) Reset(key string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.timestamps, key)
}
