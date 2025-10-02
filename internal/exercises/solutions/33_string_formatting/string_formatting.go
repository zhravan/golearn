package string_formatting

import "fmt"

func FormatName() string {
	return fmt.Sprintf("Name: %s", "\"John\"")
}

func FormatAge() string {
	return fmt.Sprintf("Age: %d", 17)
}

func FormatGpa() string {
	return fmt.Sprintf("GPA: %.2f", 3.75)
}
