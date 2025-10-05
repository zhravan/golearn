package sorting_by_functions

import "sort"

// Person represents a person with a name and age
type Person struct {
	Name string
	Age  int
}

// ByName is a type for sorting Person slices by name
type ByName []Person

// Implement sort.Interface for ByName
func (b ByName) Len() int {
	return len(b)
}

func (b ByName) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByName) Less(i, j int) bool {
	return b[i].Name < b[j].Name
}

// ByAge is a type for sorting Person slices by age
type ByAge []Person

// Implement sort.Interface for ByAge
func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

// SortByName sorts a slice of Person by name using the custom ByName type
func SortByName(people []Person) []Person {
	// Create a copy of the slice to avoid modifying the original
	result := make([]Person, len(people))
	copy(result, people)

	// Sort the copy by name
	sort.Sort(ByName(result))

	return result
}

// SortByAge sorts a slice of Person by age using the custom ByAge type
func SortByAge(people []Person) []Person {
	// Create a copy of the slice to avoid modifying the original
	result := make([]Person, len(people))
	copy(result, people)

	// Sort the copy by age
	sort.Sort(ByAge(result))

	return result
}
