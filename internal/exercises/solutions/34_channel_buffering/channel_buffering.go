package channel_buffering

import (
	"fmt"
	"slices"
	"time"
)

func optimizeImage(results chan int, image int) int {
	fmt.Println("Optimizing image...")
	time.Sleep(100 * time.Microsecond)
	results <- image + 1
	return image
}

func BatchProcessor() []int {
	images := make([]int, 100)
	for i := range 100 {
		images[i] = i
	}

	results := make(chan int, 100)
	for i := range 100 {
		go optimizeImage(results, images[i])
	}

	optimized := make([]int, 100)
	for i := range 100 {
		optimized[i] = <-results
	}

	slices.Sort(optimized)
	return optimized
}
