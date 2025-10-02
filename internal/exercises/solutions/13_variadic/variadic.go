package variadic

// Sum returns the sum of any number of integers
func Sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
