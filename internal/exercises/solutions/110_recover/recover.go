package recover_exercise

// DoWork simulates a function that might panic if the input is negative.
//
// It is designed to be called by Run, which should demonstrate
// how to safely handle the panic using recover().
func DoWork(n int) {
	if n < 0 {
		// A panic occurs if n is negative
		panic("input cannot be negative")
	}
	// Normal, non-panicking code path...
}

// Run calls DoWork and uses defer/recover to safely handle any panic.
// It returns the recovered panic value or nil if no panic occurred.
func Run(n int) (recoveredValue interface{}) {
	// The deferred function is executed just before Run returns.
	defer func() {
		// Call recover() to check if a panic has occurred.
		if r := recover(); r != nil {
			// If recover() returns a non-nil value (a panic occurred),
			// assign it to the named return variable 'recoveredValue'.
			recoveredValue = r
		}
	}()

	DoWork(n)

	// 'recoveredValue' is returned. It will be the panic value if
	// a panic occurred and was recovered, or nil otherwise.
	return recoveredValue
}
