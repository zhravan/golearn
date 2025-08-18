package arrays

import "testing"

func TestSum(t *testing.T) {
	got := Sum([5]int{1, 2, 3, 4, 5})
	if got != 15 {
		t.Fatalf("Sum(...) = %d, want 15", got)
	}
}
