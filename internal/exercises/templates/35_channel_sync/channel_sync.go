package channel_sync

import (
	"fmt"
	"time"
)

// Use case:
// Assume case where your there is an application that needs to be updated
// and you need to wait for the update to complete before shutting it down.

// TODO: 1. Return updateComplete to main go routine using a channel
// instead of printing it
func updateApp() {
	fmt.Println("Updating application...")
	time.Sleep(100 * time.Microsecond)
	updateComplete := true
	fmt.Printf("Update complete: %t", updateComplete)
}

// TODO: 2. Wait till update completes and return shutdownComplete
func CloseApp() bool {
	go updateApp()
	updateComplete := false
	shutdownComplete := updateComplete
	return shutdownComplete
}
