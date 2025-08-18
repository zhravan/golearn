package constants

import "testing"

func TestCircleArea(t *testing.T) {
	got := CircleArea(2)
	want := 12.566370614359172
	if (got-want) > 1e-9 || (want-got) > 1e-9 {
		t.Fatalf("CircleArea(2) = %v, want %v", got, want)
	}
}
