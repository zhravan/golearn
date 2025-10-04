package forloop

// SumTo returns the sum of all integers from 1 to n.
func SumTo(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}
