package range_built_in

import "testing"

func TestSumSlice(t *testing.T) {
	s := []int{1, 2, 3}
	if sumSlice(s) != 6 {
		t.Errorf("Expected 6, got %d", sumSlice(s))
	}
}

func TestCountMap(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	if countMap(m) != 2 {
		t.Errorf("Expected 2, got %d", countMap(m))
	}
}
