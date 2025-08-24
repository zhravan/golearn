package errors

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	res, err := divide(10, 2)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if res != 5 {
		t.Errorf("Expected 5, got %d", res)
	}

	res, err = divide(10, 0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if !errors.Is(err, errors.New("division by zero")) {
		t.Errorf("Expected division by zero error, got %v", err)
	}
}
