package rate_limiting

import (
	"sync"
	"time"
)

// TODO:
// - Implement a simple RateLimiter that allows N requests per interval.
// - Use a map[string][]time.Time to track request timestamps per key.
// - Use a mutex to handle concurrent access safely.
// - Implement Allow(key string) bool to check if a request is allowed.
// - Implement Reset(key string) to clear a key's request history.
type RateLimiter struct {
	mu         sync.Mutex
	limit      int
	interval   time.Duration
	timestamps map[string][]time.Time
}

// NewRateLimiter returns a new RateLimiter with the given limit and interval.
// If limit is <= 0 or interval is <= 0, it returns an error.
// The timestamps map is initialized so learners can safely implement Allow/Reset.
func NewRateLimiter(limit int, interval time.Duration) (*RateLimiter, error) {
	return &RateLimiter{
		limit:      limit,
		interval:   interval,
		timestamps: make(map[string][]time.Time),
	}, nil
}

// Allow checks if a request for the given key is allowed.
// TODO: implement logic to remove expired timestamps and enforce limit.
func (r *RateLimiter) Allow(key string) bool {
	// TODO: remove expired timestamps for this key
	// TODO: allow the request if below limit, otherwise deny
	return false
}

// Reset clears the request history for a given key.
// If the key does not exist, it does nothing.
func (r *RateLimiter) Reset(key string) {
	// TODO: implement reset logic
}
