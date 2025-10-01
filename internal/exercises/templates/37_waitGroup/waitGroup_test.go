package waitgroup

import "testing"

func TestHello(t *testing.T) {
	got := waitGroup()
	want := "Worker done"
	if got != want {
		t.Fatalf("waitGroup() = %q, want %q", got, want)
	}
}
