package mutexes

import "testing"

func TestCounter(t *testing.T) {
	got := Counting()
	want := numWorkers

	if got != want {
		t.Fatalf("Expected %d, got %d", want, got)
	}
}
