package sorting

import (
	"slices"
)

var Years = []int{2017, 2003, 2026}
var Pets = []string{"Dog", "Cat", "Parrot"}

func SortNumbers() []int {
	slices.Sort(Years)
	return Years
}

func SortAnimals() []string {
	slices.Sort(Pets)
	return Pets
}
