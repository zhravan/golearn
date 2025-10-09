package sorting

import (
	"slices"
	"testing"
)

func TestSorting(t *testing.T) {
	t.Run("Sorting numbers", func(t *testing.T) {
		SortYears()
		if !slices.IsSorted(Years) {
			t.Fatal("Years are not sorted!")
		}
	})

	t.Run("Sorting strings", func(t *testing.T) {
		SortPets()
		if !slices.IsSorted(Pets) {
			t.Fatal("Pets are not sorted!")
		}
	})
}
