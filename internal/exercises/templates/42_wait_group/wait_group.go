package waitgroup

import "sync"

// Squares takes a slice of numbers and should return a slice
// containing the square of each number.

// Use goroutines to perform the calculations concurrently,
// and use sync.WaitGroup to make sure all goroutines finish
// before returning the results.

// Steps:
// 1. Create a channel to collect results.
// 2. Start one goroutine per number.
// 3. Each goroutine should send n*n into the channel.
// 4. Wait for all goroutines to finish, then close the channel.
// 5. Collect all values from the channel into a slice and return it.

func Squares(nums []int) []int {
	var wg sync.WaitGroup

	// TODO: implement the logic here

	return nil
}
