package atomic_counters

import "testing"

func TestAtomicCounter(t *testing.T) {

	t.Run("Counter values is correct!", func(t *testing.T) {
		got := IncrementCounter()
		want := 10_000
		if int(got) != want {
			t.Fatalf("Got %d, want %d", got, want)
		}
	})
}
