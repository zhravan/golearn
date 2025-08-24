package generics

func identity[T any](arg T) T {
	return arg
}

func sumIntsOrFloats[T int | float64](a, b T) T {
	return a + b
}
