package variadic

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(1, 2, 3, 4); got != 10 {
		t.Fatalf("Sum(1,2,3,4) = %d, want 10", got)
	}
}
