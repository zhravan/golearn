package rate_limiting

import (
	"testing"
	"time"
)

func TestRateLimiter(t *testing.T) {
	rl := NewRateLimiter(2, 100*time.Millisecond)
	key := "user1"

	if !rl.Allow(key) {
		t.Error("Expected first request to be allowed")
	}
	if !rl.Allow(key) {
		t.Error("Expected second request to be allowed")
	}
	if rl.Allow(key) {
		t.Error("Expected third request to be denied")
	}

	// Wait for interval to pass
	time.Sleep(110 * time.Millisecond)

	if !rl.Allow(key) {
		t.Error("Expected request after interval to be allowed")
	}
}

func TestReset(t *testing.T) {
	rl := NewRateLimiter(1, 1*time.Second)
	key := "user2"

	if !rl.Allow(key) {
		t.Error("Expected first request to be allowed")
	}

	err := rl.Reset(key)
	if err != nil {
		t.Errorf("Unexpected error on reset: %v", err)
	}

	if !rl.Allow(key) {
		t.Error("Expected request to be allowed after reset")
	}
}