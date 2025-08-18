package functions

import "testing"

func TestApply(t *testing.T) {
	dbl := func(n int) int { return n * 2 }
	if got := Apply(dbl, 7); got != 14 {
		t.Fatalf("Apply(dbl, 7) = %d, want 14", got)
	}
}
