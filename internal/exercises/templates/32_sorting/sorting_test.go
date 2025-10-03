package sorting

import (
	"slices"
	"testing"
)

func TestSorting(t *testing.T) {
	t.Run("Sorting numbers", func(t *testing.T) {
		SortNumbers()
		if !slices.IsSorted(Years) {
			t.Fatal("Strings are not sorted!")
		}
	})

	t.Run("Sorting strings", func(t *testing.T) {
		SortAnimals()
		if !slices.IsSorted(Pets) {
			t.Fatal("Strings are not sorted!")
		}
	})
}
