package channel_buffering

import (
	"fmt"
	"slices"
	"time"
)

// Use case: Batch processing
// Assume you are optimizing images in a batch.
// You have a function optimizeImage that optimize an images.
// You want to process 100 images in parallel.

// TODO: 1. Use a channel to buffer 100 image optimizations
func optimizeImage(image int) int {
	fmt.Println("Optimizing image...")
	time.Sleep(100 * time.Microsecond)
	return image + 1
}

// TODO: 2. Wait till all 100 images are optimized
func BatchProcessor() []int {
	images := make([]int, 100)
	for i := range 100 {
		images[i] = i
	}

	optimized := make([]int, 100)
	for i := range 100 {
		optimized[i] = optimizeImage(images[i])
	}

	slices.Sort(optimized)
	return optimized
}
