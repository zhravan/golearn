package multireturn

import "testing"

func TestDivMod(t *testing.T) {
	q, r, err := DivMod(7, 3)
	if err != nil || q != 2 || r != 1 {
		t.Fatalf("DivMod(7,3) = (%d,%d,%v), want (2,1,nil)", q, r, err)
	}
	_, _, err = DivMod(1, 0)
	if err == nil {
		t.Fatalf("expected error for divide by zero")
	}
}
