package waitgroup

import (
	"reflect"
	"sort"
	"testing"
	"time"
)

func equalIgnoreOrder(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	x, y := append([]int(nil), a...), append([]int(nil), b...)
	sort.Ints(x)
	sort.Ints(y)
	return reflect.DeepEqual(x, y)
}

func TestSquaresBasic(t *testing.T) {
	got := Squares([]int{1, 2, 3, 4})
	want := []int{1, 4, 9, 16}
	if !equalIgnoreOrder(got, want) {
		t.Fatalf("Squares() = %v; want %v", got, want)
	}
}

func TestSquaresEmpty(t *testing.T) {
	if len(Squares([]int{})) != 0 {
		t.Fatal("expected empty result for empty input")
	}
}

func TestSquaresTimeout(t *testing.T) {
	done := make(chan struct{})
	go func() {
		_ = Squares([]int{5, 6, 7})
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(2 * time.Second):
		t.Fatal("function deadlocked")
	}
}
