package channel_sync

import (
	"fmt"
	"time"
)

func updateApp(done chan bool) {
	fmt.Println("Updating application...")
	time.Sleep(100 * time.Microsecond)
	updateComplete := true
	done <- updateComplete
}

func CloseApp() bool {
	done := make(chan bool, 1)
	go updateApp(done)
	updateComplete := <-done
	shutdownComplete := updateComplete
	return shutdownComplete
}
