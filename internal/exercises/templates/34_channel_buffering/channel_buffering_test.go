package channel_buffering

import (
	"reflect"
	"testing"
	"time"
)

func TestChannelBuffering(t *testing.T) {
	t.Run("Images were optimized successfully!", func(t *testing.T) {
		got := BatchProcessor()
		want := make([]int, 100)
		for i := range 100 {
			want[i] = i + 1
		}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("Expected %v, got %v", want, got)
		}
	})

	t.Run("Image optimization was fast!", func(t *testing.T) {
		start := time.Now()
		BatchProcessor()
		elapsed := time.Since(start)

		if elapsed.Milliseconds() >= 5 {
			t.Fatal("Optimizing took too long!")
		}
	})
}

// For benchmarking BatchProcessor to set avg value in test
func BenchmarkCloseApp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BatchProcessor()
	}
}
