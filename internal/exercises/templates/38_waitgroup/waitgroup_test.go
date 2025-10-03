package waitgroup

import "testing"

func TestWaitGroup(t *testing.T) {
	got := waitGroup()
	want := "Worker done"
	if got != want {
		t.Fatalf("waitGroup() = %q, want %q", got, want)
	}
}
