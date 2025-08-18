package forloop

import "testing"

func TestSumTo(t *testing.T) {
	if got := SumTo(10); got != 55 {
		t.Fatalf("SumTo(10) = %d, want 55", got)
	}
}
