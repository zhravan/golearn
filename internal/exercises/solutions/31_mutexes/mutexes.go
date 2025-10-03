package mutexes

import "sync"

const numWorkers = 10_000

func Counting() int {

	count := 0
	done := make(chan bool, numWorkers)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			count += 1
			mu.Unlock()
			done <- true
		}()
	}

	for i := 1; i <= numWorkers; i++ {
		<-done
	}

	wg.Wait()
	return count
}
