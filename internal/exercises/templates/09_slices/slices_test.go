package slices

import "testing"

func TestAppendAndSum(t *testing.T) {
	got := AppendAndSum([]int{1, 2, 3}, 4)
	if got != 10 {
		t.Fatalf("AppendAndSum(...) = %d, want 10", got)
	}
}
