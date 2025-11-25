package range_over_channels

import (
	"testing"
)

func TestRangeOverChannels(t *testing.T) {
	got := processStream()
	want := 45

	if got != want {
		t.Fatalf("Expected %v, got %v", want, got)
	}
}
