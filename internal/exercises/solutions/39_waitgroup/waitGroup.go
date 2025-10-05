package waitgroup

import (
	"sync"
)

// The waitGroup in Go is used to wait for a collection of goroutines to finish executing

func worker(wg *sync.WaitGroup, result *string) {
	defer wg.Done()
	*result = "Worker done"
}

func waitGroup() string {
	var wg sync.WaitGroup
	result := ""
	wg.Add(1)
	go worker(&wg, &result)
	wg.Wait()
	return result
}
