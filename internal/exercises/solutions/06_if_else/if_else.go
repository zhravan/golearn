package ifelse

// Sign returns "negative", "zero", or "positive" based on n
func Sign(n int) string {
	if n < 0 {
		return "negative"
	} else if n == 0 {
		return "zero"
	} else {
		return "positive"
	}
}
