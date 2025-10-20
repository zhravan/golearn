package rate_limiting

import (
	"errors"
	"sync"
	"time"
)

// TODO:
// - Implement a simple RateLimiter that allows N requests per interval.
// - Use a map[string][]time.Time to track timestamps of requests per key.
// - Use a mutex to handle concurrent access safely.
// - Implement Allow(key string) bool method returning true if allowed, false otherwise.
// - Optionally, implement Reset(key string) error to clear a key's request history.
type RateLimiter struct {
	mu         sync.Mutex
	limit      int
	interval   time.Duration
	timestamps map[string][]time.Time
}

// NewRateLimiter should return a pointer to a new RateLimiter with initialized fields.
func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	// TODO: initialize and return RateLimiter
	return &RateLimiter{}
}

// Allow returns true if a request for the given key is allowed, false if the limit is exceeded.
func (r *RateLimiter) Allow(key string) bool {
	// TODO: implement rate limiting logic
	return false
}

// Reset clears the request history for a given key.
func (r *RateLimiter) Reset(key string) error {
	// TODO: implement reset logic
	return errors.New("not implemented")
}