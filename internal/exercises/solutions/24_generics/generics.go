package generics

// identity returns the argument, of any type, as is.
func identity[T any](arg T) T {
	return arg
}

// sumIntsOrFloats takes two arguments of the same type, either int or float64, and returns their sum.
func sumIntsOrFloats[T int | float64](a, b T) T {
	return a + b
}
