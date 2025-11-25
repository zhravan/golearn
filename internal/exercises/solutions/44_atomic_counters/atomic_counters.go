package atomic_counters

import (
	"sync"
	"sync/atomic"
)

func NoRequestsProcessed() uint64 {

	var ops atomic.Uint64
	var wg sync.WaitGroup

	wg.Add(10_000)
	for range 10_000 {
		go func() {
			defer wg.Done()
			ops.Add(1)
		}()
	}
	wg.Wait()
	return ops.Load()
}
