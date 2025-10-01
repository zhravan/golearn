package recover_exercise

// DoWork simulates a function that might panic if the input is negative.
func DoWork(n int) {
	if n < 0 {
		// A panic occurs if n is negative
		panic("input cannot be negative")
	}
	// Normal, non-panicking code path...
}

// Run should call DoWork and safely recover from any panic
// that occurs during its execution, returning the recovered value.
// It should return nil if no panic occurred.
func Run(n int) (recoveredValue interface{}) {
	// 1. Add your 'defer' statement here.
	// 2. The deferred function should call recover() and assign the result
	//    to 'recoveredValue' if it is not nil.

	// Your code here:

	DoWork(n)

	return recoveredValue
}