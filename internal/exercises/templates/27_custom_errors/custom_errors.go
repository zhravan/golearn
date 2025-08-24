package custom_errors

import (
	"fmt"
)

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error Code %d: %s", e.Code, e.Message)
}

func processInput(input int) (string, *MyError) {
	if input < 0 {
		return "", &MyError{Code: 1001, Message: "Input cannot be negative"}
	}
	return "Processing complete", nil
}
