package waitgroup

import (
	"sync"
)

// TODO: Implement these functions so tests pass

func worker(wg *sync.WaitGroup, result *string) {
	*result = "Worker done"
}

func waitGroup() string {
	var wg sync.WaitGroup
	result := ""
	go worker(&wg, &result)
	return result
}
