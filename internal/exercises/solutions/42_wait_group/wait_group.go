package waitgroup

import "sync"

func Squares(nums []int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}

	ch := make(chan int, length)
	var wg sync.WaitGroup
	wg.Add(length)

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

	out := make([]int, 0, length)
	for v := range ch {
		out = append(out, v)
	}
	return out
}
