package string_formatting

import "fmt"

func FormatName() string {
	// TODO: format the name as a string
	return fmt.Sprintf("Name: ", "\"John\"")
}

func FormatAge() string {
	// TODO: format the age as a digit
	return fmt.Sprintf("Age: ", 17)
}

func FormatGpa() string {
	// TODO: format the GPA for floating point number with 2 decimal places
	return fmt.Sprintf("GPA: ", 3.75)
}
