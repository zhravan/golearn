package channel_sync

import (
	"testing"
	"time"
)

func TestChannelSync(t *testing.T) {

	t.Run("Wait till update completes", func(t *testing.T) {
		start := time.Now()
		CloseApp()
		elapsed := time.Since(start)

		if elapsed.Microseconds() <= 100 {
			t.Fatal("Update incomplete!")
		}
	})

	t.Run("Update completed successfully", func(t *testing.T) {
		got := CloseApp()
		want := true
		if got != want {
			t.Fatal("Update unsuccessful!")
		}
	})

}

// For benchmarking CloseApp to set avg value in test
func BenchmarkCloseApp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CloseApp()
	}
}
