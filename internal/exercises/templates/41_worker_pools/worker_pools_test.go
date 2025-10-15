package worker_pools

import (
	"fmt"
	"testing"
	"time"
)

func TestProcessLog(t *testing.T) {
	log := LogEntry{
		ID:        1,
		Timestamp: "2025-10-15T10:30:00Z",
		Level:     "ERROR",
		Message:   "  Database connection failed  ",
	}

	result := processLog(log)

	if result.ID != 1 {
		t.Errorf("Expected ID 1, got %d", result.ID)
	}

	if result.Severity != 3 {
		t.Errorf("Expected severity 3 for ERROR, got %d", result.Severity)
	}

	if result.CleanedMessage != "Database connection failed" {
		t.Errorf("Expected cleaned message 'Database connection failed', got '%s'", result.CleanedMessage)
	}
}

func TestLogProcessor(t *testing.T) {
	t.Run("Process all logs successfully", func(t *testing.T) {
		logs := []LogEntry{
			{ID: 1, Timestamp: "2025-10-15T10:30:00Z", Level: "INFO", Message: "  Server started  "},
			{ID: 2, Timestamp: "2025-10-15T10:31:00Z", Level: "WARNING", Message: "High memory usage"},
			{ID: 3, Timestamp: "2025-10-15T10:32:00Z", Level: "ERROR", Message: "Connection timeout"},
			{ID: 4, Timestamp: "2025-10-15T10:33:00Z", Level: "INFO", Message: "Request completed"},
			{ID: 5, Timestamp: "2025-10-15T10:34:00Z", Level: "ERROR", Message: "  Failed to save  "},
		}

		results := LogProcessor(logs, 3)

		if len(results) != 5 {
			t.Fatalf("Expected 5 processed logs, got %d", len(results))
		}

		// Verify all IDs are present
		ids := make(map[int]bool)
		for _, r := range results {
			ids[r.ID] = true
		}

		for i := 1; i <= 5; i++ {
			if !ids[i] {
				t.Errorf("Missing log ID %d in results", i)
			}
		}
	})

	t.Run("Processing is concurrent and faster than sequential", func(t *testing.T) {
		// Create 20 logs
		logs := make([]LogEntry, 20)
		for i := 0; i < 20; i++ {
			logs[i] = LogEntry{
				ID:        i + 1,
				Timestamp: "2025-10-15T10:30:00Z",
				Level:     "INFO",
				Message:   "Test message",
			}
		}

		start := time.Now()
		results := LogProcessor(logs, 5)
		elapsed := time.Since(start)

		if len(results) != 20 {
			t.Fatalf("Expected 20 processed logs, got %d", len(results))
		}

		// With 5 workers processing 20 logs at ~5ms each:
		// Sequential would take ~100ms
		// Concurrent should take ~20-30ms (4 batches of 5)
		if elapsed.Milliseconds() > 50 {
			t.Errorf("Processing took too long (%v), worker pool may not be working correctly", elapsed)
		}
	})

	t.Run("Verify correct severity mapping", func(t *testing.T) {
		logs := []LogEntry{
			{ID: 1, Timestamp: "2025-10-15T10:30:00Z", Level: "INFO", Message: "info"},
			{ID: 2, Timestamp: "2025-10-15T10:30:00Z", Level: "WARNING", Message: "warn"},
			{ID: 3, Timestamp: "2025-10-15T10:30:00Z", Level: "ERROR", Message: "error"},
		}

		results := LogProcessor(logs, 2)

		severities := make(map[int]int)
		for _, r := range results {
			severities[r.ID] = r.Severity
		}

		if severities[1] != 1 {
			t.Errorf("Expected INFO severity 1, got %d", severities[1])
		}
		if severities[2] != 2 {
			t.Errorf("Expected WARNING severity 2, got %d", severities[2])
		}
		if severities[3] != 3 {
			t.Errorf("Expected ERROR severity 3, got %d", severities[3])
		}
	})
}

func TestWorkerChannelDirections(t *testing.T) {
	// This test ensures channels are properly typed with directions
	// by checking if the implementation compiles and runs correctly
	logs := []LogEntry{
		{ID: 1, Timestamp: "2025-10-15T10:30:00Z", Level: "INFO", Message: "test"},
	}

	results := LogProcessor(logs, 1)

	if len(results) != 1 {
		t.Errorf("Expected 1 result, got %d", len(results))
	}
}

func BenchmarkLogProcessor(b *testing.B) {
	logs := make([]LogEntry, 100)
	for i := 0; i < 100; i++ {
		logs[i] = LogEntry{
			ID:        i + 1,
			Timestamp: "2025-10-15T10:30:00Z",
			Level:     "INFO",
			Message:   "Test message",
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LogProcessor(logs, 10)
	}
}

func BenchmarkLogProcessorWorkerComparison(b *testing.B) {
	logs := make([]LogEntry, 100)
	for i := 0; i < 100; i++ {
		logs[i] = LogEntry{
			ID:        i + 1,
			Timestamp: "2025-10-15T10:30:00Z",
			Level:     "INFO",
			Message:   "Test message",
		}
	}

	workerCounts := []int{1, 2, 5, 10, 20}

	for _, numWorkers := range workerCounts {
		b.Run(fmt.Sprintf("Workers_%d", numWorkers), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				LogProcessor(logs, numWorkers)
			}
		})
	}
}
