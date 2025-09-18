package errors

import (
	"fmt"
	"os"
)

// divZeroErr is a minimal error type whose message is "division by zero".
type divZeroErr struct{}

// Error makes divZeroErr implement error.
func (divZeroErr) Error() string { return "division by zero" }

// Is allows errors.Is(err, errors.New("division by zero")) to succeed.
func (divZeroErr) Is(target error) bool {
	if target == nil {
		return false
	}
	return target.Error() == "division by zero"
}

// divide returns a/b or a "division by zero" error when b == 0.
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, divZeroErr{}
	}
	return a / b, nil
}

// processDivision prints the result or the error to stdout.
func processDivision(a, b int) {
	res, err := divide(a, b)
	if err != nil {
		fmt.Fprint(os.Stdout, "Error: ", err, "\n")
		return
	}
	fmt.Fprint(os.Stdout, "Result: ", res, "\n")
}
