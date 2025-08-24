package custom_errors

import "testing"

func TestProcessInput(t *testing.T) {
	_, err := processInput(-1)
	if err == nil {
		t.Errorf("Expected an error for negative input, got nil")
	}

	if err.Code != 1001 {
		t.Errorf("Expected error code 1001, got %d", err.Code)
	}

	if err.Message != "Input cannot be negative" {
		t.Errorf("Expected error message 'Input cannot be negative', got %s", err.Message)
	}

	_, err = processInput(10)
	if err != nil {
		t.Errorf("Expected no error for positive input, got %v", err)
	}
}
