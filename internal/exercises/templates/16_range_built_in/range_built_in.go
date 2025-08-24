package range_built_in

func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func countMap(m map[string]int) int {
	count := 0
	for _, _ = range m {
		count++
	}
	return count
}
