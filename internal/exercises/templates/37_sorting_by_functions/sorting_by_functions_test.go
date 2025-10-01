package sorting_by_functions

import (
	"reflect"
	"testing"
)

func TestSortByName(t *testing.T) {
	people := []Person{
		{Name: "Charlie", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}

	expected := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
		{Name: "Charlie", Age: 30},
	}

	result := SortByName(people)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortByName() = %v, want %v", result, expected)
	}

	// Ensure original slice is not modified
	originalExpected := []Person{
		{Name: "Charlie", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}
	if !reflect.DeepEqual(people, originalExpected) {
		t.Errorf("Original slice was modified: %v, want %v", people, originalExpected)
	}
}

func TestSortByAge(t *testing.T) {
	people := []Person{
		{Name: "Charlie", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}

	expected := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Charlie", Age: 30},
		{Name: "Bob", Age: 35},
	}

	result := SortByAge(people)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("SortByAge() = %v, want %v", result, expected)
	}

	// Ensure original slice is not modified
	originalExpected := []Person{
		{Name: "Charlie", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}
	if !reflect.DeepEqual(people, originalExpected) {
		t.Errorf("Original slice was modified: %v, want %v", people, originalExpected)
	}
}

func TestByNameInterface(t *testing.T) {
	people := []Person{
		{Name: "Charlie", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}

	byName := ByName(people)

	// Test Len
	if byName.Len() != 3 {
		t.Errorf("ByName.Len() = %d, want 3", byName.Len())
	}

	// Test Less
	if !byName.Less(1, 0) { // "Alice" < "Charlie"
		t.Errorf("ByName.Less(1, 0) should be true (Alice < Charlie)")
	}
	if byName.Less(0, 1) { // "Charlie" > "Alice"
		t.Errorf("ByName.Less(0, 1) should be false (Charlie > Alice)")
	}

	// Test Swap
	original := make([]Person, len(people))
	copy(original, people)
	byName.Swap(0, 1)
	if byName[0].Name != original[1].Name || byName[1].Name != original[0].Name {
		t.Errorf("ByName.Swap(0, 1) did not swap correctly")
	}
}

func TestByAgeInterface(t *testing.T) {
	people := []Person{
		{Name: "Charlie", Age: 30},
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 35},
	}

	byAge := ByAge(people)

	// Test Len
	if byAge.Len() != 3 {
		t.Errorf("ByAge.Len() = %d, want 3", byAge.Len())
	}

	// Test Less
	if !byAge.Less(1, 0) { // 25 < 30
		t.Errorf("ByAge.Less(1, 0) should be true (25 < 30)")
	}
	if byAge.Less(0, 1) { // 30 > 25
		t.Errorf("ByAge.Less(0, 1) should be false (30 > 25)")
	}

	// Test Swap
	original := make([]Person, len(people))
	copy(original, people)
	byAge.Swap(0, 1)
	if byAge[0].Age != original[1].Age || byAge[1].Age != original[0].Age {
		t.Errorf("ByAge.Swap(0, 1) did not swap correctly")
	}
}