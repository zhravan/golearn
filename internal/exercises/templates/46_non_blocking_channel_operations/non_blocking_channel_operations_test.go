package non_blocking_channel_operations

import (
	"testing"
	"time"
)

func TestNonBlockingChannelOperations(t *testing.T) {

	t.Run("Logs are processed without dropping", func(t *testing.T) {

		processingDelay = 20 * time.Microsecond

		logger := NewLogger(10)

		for range 10 {
			logger.Log("msg")
		}
		logger.Close()

		processed, dropped := logger.Stats()

		if dropped != 0 {
			t.Fatalf("Expected no dropped logs, got %d", dropped)
		}
		if processed != 10 {
			t.Fatalf("Expected 10 processed logs, got %d", processed)
		}
	})

	t.Run("Logs are not dropped", func(t *testing.T) {
		logger := NewLogger(5)

		for range 10 {
			logger.Log("msg")
		}

		logger.Close()

		processed, dropped := logger.Stats()

		if dropped != 5 {
			t.Fatalf("Expected no dropped logs, got %d", dropped)
		}
		if processed != 5 {
			t.Fatalf("Expected 10 processed logs, got %d", processed)
		}
	})
}
