package rate_limiting

import (
	"sync"
	"sync/atomic"
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

	// Verify limit is reached before reset
	if rl.Allow(key) {
		t.Error("Expected third request to be denied")
	}

	// Wait for interval to pass with buffer
	time.Sleep(150 * time.Millisecond)

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

	// Verify limit is reached before reset
	if rl.Allow(key) {
		t.Error("Expected second request to be denied before reset")
	}

	rl.Reset(key)

	if !rl.Allow(key) {
		t.Error("Expected request to be allowed after reset")
	}
}

func TestRateLimiterConcurrent(t *testing.T) {
	rl := NewRateLimiter(10, 100*time.Millisecond)
	key := "user1"

	var wg sync.WaitGroup
	var allowed atomic.Int32

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if rl.Allow(key) {
				allowed.Add(1)
			}
		}()
	}

	wg.Wait()
	if allowed.Load() != 10 {
		t.Errorf("Expected 10 allowed requests, got %d", allowed.Load())
	}
}

func TestMultipleKeys(t *testing.T) {
	rl := NewRateLimiter(1, 100*time.Millisecond)

	if !rl.Allow("userA") {
		t.Error("Expected first request for userA to be allowed")
	}
	if !rl.Allow("userB") {
		t.Error("Expected first request for userB to be allowed")
	}
	if rl.Allow("userA") {
		t.Error("Expected second request for userA to be denied")
	}
	if rl.Allow("userB") {
		t.Error("Expected second request for userB to be denied")
	}
}