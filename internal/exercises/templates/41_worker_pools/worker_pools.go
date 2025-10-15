package worker_pools

import (
	"fmt"
	"time"
)

// Use case: Log Processing System
// You're building a log processing system that needs to analyze logs from
// multiple servers. Each log entry needs to be parsed, validated, and
// transformed before being stored. Using a worker pool pattern ensures
// efficient processing of thousands of log entries without overwhelming
// system resources.

// LogEntry represents a single log entry to process
type LogEntry struct {
	ID        int
	Timestamp string
	Level     string
	Message   string
}

// ProcessedLog represents a log after processing
type ProcessedLog struct {
	ID             int
	Severity       int // 1=INFO, 2=WARNING, 3=ERROR
	CleanedMessage string
}

// TODO: 1. Implement processLog function that processes a single log entry
// It should:
// - Convert log level string to severity int (INFO=1, WARNING=2, ERROR=3)
// - Clean the message by trimming spaces
// - Return a ProcessedLog
func processLog(log LogEntry) ProcessedLog {
	fmt.Printf("Processing log ID: %d\n", log.ID)
	time.Sleep(5 * time.Millisecond) // Simulate processing time

	// TODO: Implement log processing logic
	return ProcessedLog{}
}

// TODO: 2. Implement worker function that:
// - Takes an ID, jobs channel (receive-only), and results channel (send-only)
// - Continuously receives LogEntry from jobs channel
// - Processes each log using processLog
// - Sends ProcessedLog to results channel
// - Stops when jobs channel is closed
func worker(id int, jobs <-chan LogEntry, results chan<- ProcessedLog) {
	// TODO: Implement worker logic with proper channel directions
	fmt.Println("Worker started")
}

// TODO: 3. Implement LogProcessor that:
// - Creates jobs and results channels
// - Spawns numWorkers worker goroutines
// - Sends all logs to the jobs channel in a separate goroutine
// - Collects all processed results
// - Returns slice of ProcessedLog
func LogProcessor(logs []LogEntry, numWorkers int) []ProcessedLog {
	// TODO: Create channels

	// TODO: Start worker pool

	// TODO: Send jobs

	// TODO: Collect results

	processed := []ProcessedLog{}
	return processed
}
