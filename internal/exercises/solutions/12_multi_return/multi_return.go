package multireturn

import "fmt"

// DivMod divides a by b, returning quotient, remainder, and error for divide-by-zero
func DivMod(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("cannot divide by zero")
	}
	quotient := a / b
	remainder := a % b
	return quotient, remainder, nil
}
