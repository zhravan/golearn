package closures

// closure returns a function that increments and returns an internal counter
func closure() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}
