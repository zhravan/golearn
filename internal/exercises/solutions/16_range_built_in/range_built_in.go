package range_built_in

// sumSlice returns the sum of all elements in the slice using range
func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// countMap returns the number of key/value pairs in the map using range
func countMap(m map[string]int) int {
	count := 0
	for range m {
		count++
	}
	return count
}
