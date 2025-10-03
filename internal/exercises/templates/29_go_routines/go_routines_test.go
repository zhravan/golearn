package go_routines

import (
	"testing"
	"time"
)

func TestGoRoutines(t *testing.T) {
	start := time.Now()
	RunConcurrently()
	elapsed := time.Since(start)

	// according to benchmark it is about ~2 microseconds
	// x2 for more breathing room, so 4
	if elapsed.Microseconds() >= 4 {
		t.Fatal("Go routines was not used!")
	}
}

func BenchmarkRunConcurrently(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunConcurrently()
	}
}
