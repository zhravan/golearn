package mutexes

import "sync"

const numWorkers = 10_000

func Counting() int {

	// TODO: update following code to avoid race conditions
	// using sync.Mutex
	count := 0
	done := make(chan bool, numWorkers)
	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count += 1
			done <- true
		}()
	}

	for i := 1; i <= numWorkers; i++ {
		<-done
	}

	wg.Wait()
	return count
}
