// regular_expressions_solution.go
// Solution for: Regular Expressions
package solution

import "regexp"

// MatchRegex returns true if the input matches the pattern, false otherwise.
func MatchRegex(pattern, input string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(input)
}
