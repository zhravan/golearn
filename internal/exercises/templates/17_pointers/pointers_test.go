package pointers

import "testing"

func TestModifyValue(t *testing.T) {
	x := 10
	ptr := &x
	newPtr := modifyValue(ptr)
	if *newPtr != 100 {
		t.Errorf("Expected 100, got %d", *newPtr)
	}
	if newPtr != ptr {
		t.Errorf("Expected same pointer, got different")
	}
}
