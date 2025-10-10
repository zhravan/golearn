package channel_select

import (
	"testing"
	"time"
)

func TestChannelSelect(t *testing.T) {

	t.Run("Network request result either success or timeout", func(t *testing.T) {

		start := time.Now()
		got := NetworkRequest()
		elapsed := time.Since(start)

		if elapsed.Microseconds() < 5 {
			if got != "Success!" {
				t.Fatalf("Request should be successful, but got: %s", got)
			}
		} else {
			if got != "Timeout!" {
				t.Fatalf("Request should be timeout, but got: %s", got)
			}
		}
	})
}
