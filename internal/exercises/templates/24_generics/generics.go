package generics

// TODO:
// - Implement a generic identity function: identity[T any](arg T) T.
// - Implement sumIntsOrFloats that accepts either ints or float64s and returns their sum.

func identity[T any](arg T) T {
	// TODO: return the argument back
	var zero T
	return zero
}

func sumIntsOrFloats[T int | float64](a, b T) T {
	// TODO: return a + b
	var zero T
	return zero
}
