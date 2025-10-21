package waitgroup

import "sync"

func Squares(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	ch := make(chan int, len(nums))
	var wg sync.WaitGroup
	wg.Add(len(nums))

	for _, n := range nums {
		n := n
		go func() {
			defer wg.Done()
			ch <- n * n
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	out := make([]int, 0, len(nums))
	for v := range ch {
		out = append(out, v)
	}
	return out
}
