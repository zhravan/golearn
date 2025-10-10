package channel_directions

import (
	"fmt"
	"slices"
	"time"
)

// Use case: Database pool.

// TODO: 5. Make sure channels parameters are type safe by specifying the
// channel directions
func worker() {
	fmt.Println("Processing query ...")
	time.Sleep(10 * time.Microsecond)

}

// TODO: 1. Create a channel to send queries to the workers
// TODO: 2. Create a channel to receive results from the workers
// TODO: 3. Create a pool of five workers and send queries to them
// TODO: 4. Receive results from the workers and store them in a rs slice
func DbPool() []int {

	worker()

	rs := []int{}

	slices.Sort(rs)
	return rs
}
