package atomic_counters

// TODO: Assume you have to keep track of the number of requests processed
// by an applications to display on a dashboard.
// 1. Launch 10_000 goroutines that increment a counter.
//  2. Wait for all goroutines to finish and return the final value of the counter.
func NoRequestsProcessed() uint64 {
	return 0
}
