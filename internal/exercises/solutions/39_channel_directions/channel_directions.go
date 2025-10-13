package channel_directions

import (
	"fmt"
	"slices"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	fmt.Printf("Worker ID: %d\n", id)
	fmt.Println("Processing query ...")
	time.Sleep(10 * time.Microsecond)
	for job := range jobs {
		results <- job * 2
	}
}

func DbPool() []int {
	jobs := make(chan int)
	results := make(chan int)

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for i := 1; i <= 50; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	rs := []int{}
	for a := 1; a <= 50; a++ {
		rs = append(rs, <-results)
	}

	slices.Sort(rs)
	return rs
}
