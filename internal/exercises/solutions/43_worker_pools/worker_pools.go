package worker_pools

import (
	"fmt"
	"strings"
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

// processLog processes a single log entry
func processLog(log LogEntry) ProcessedLog {
	fmt.Printf("Processing log ID: %d\n", log.ID)
	time.Sleep(5 * time.Microsecond) // Simulate processing time

	// Convert level to severity
	severity := 1 // Default INFO
	switch log.Level {
	case "INFO":
		severity = 1
	case "WARNING":
		severity = 2
	case "ERROR":
		severity = 3
	}

	// Clean message
	cleanedMessage := strings.TrimSpace(log.Message)

	return ProcessedLog{
		ID:             log.ID,
		Severity:       severity,
		CleanedMessage: cleanedMessage,
	}
}

// worker receives jobs from a channel, processes them, and sends results
func worker(id int, jobs <-chan LogEntry, results chan<- ProcessedLog) {
	for job := range jobs {
		fmt.Printf("Worker %d processing log %d\n", id, job.ID)
		processed := processLog(job)
		results <- processed
	}
}

// LogProcessor creates a worker pool and processes all logs
func LogProcessor(logs []LogEntry, numWorkers int) []ProcessedLog {
	// Create channels
	jobs := make(chan LogEntry, len(logs))
	results := make(chan ProcessedLog, len(logs))

	// Start worker pool
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs
	go func() {
		for _, log := range logs {
			jobs <- log
		}
		close(jobs)
	}()

	// Collect results
	processed := make([]ProcessedLog, 0, len(logs))
	for i := 0; i < len(logs); i++ {
		result := <-results
		processed = append(processed, result)
	}

	return processed
}
