package custom_errors


// TODO:
// - Define a custom error type that implements the error interface.
// - Include a Code and Message, and format them in Error().
// - Implement processInput to return an error for invalid input (negative).

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	// TODO: format the error message
	return ""
}

func processInput(input int) (string, *MyError) {
	// TODO: return an error for negative input, otherwise success
	return "", nil
}
