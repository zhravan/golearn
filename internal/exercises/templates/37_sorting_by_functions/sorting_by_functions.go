package sorting_by_functions

import "sort"

// TODO:
// - Complete the Person struct with Name and Age fields
// - Implement the sort.Interface methods for ByName and ByAge
// - Complete SortByName and SortByAge functions to sort slices using custom comparison

// Person represents a person with a name and age
type Person struct {
	Name string
	Age  int
}

// ByName is a type for sorting Person slices by name
type ByName []Person

// Implement sort.Interface for ByName
func (b ByName) Len() int {
	// TODO: return the length of the slice
	return 0
}

func (b ByName) Swap(i, j int) {
	// TODO: swap elements at positions i and j
}

func (b ByName) Less(i, j int) bool {
	// TODO: return true if element at i should come before element at j (compare by name)
	return false
}

// ByAge is a type for sorting Person slices by age
type ByAge []Person

// Implement sort.Interface for ByAge
func (b ByAge) Len() int {
	// TODO: return the length of the slice
	return 0
}

func (b ByAge) Swap(i, j int) {
	// TODO: swap elements at positions i and j
}

func (b ByAge) Less(i, j int) bool {
	// TODO: return true if element at i should come before element at j (compare by age)
	return false
}

// SortByName sorts a slice of Person by name using the custom ByName type
func SortByName(people []Person) []Person {
	// TODO: create a copy of the slice and sort it by name
	// Hint: use sort.Sort() with ByName type
	return nil
}

// SortByAge sorts a slice of Person by age using the custom ByAge type
func SortByAge(people []Person) []Person {
	// TODO: create a copy of the slice and sort it by age
	// Hint: use sort.Sort() with ByAge type
	return nil
}