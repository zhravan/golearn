package methods

import "testing"

func TestArea(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	if r.Area() != 50 {
		t.Errorf("Expected 50, got %d", r.Area())
	}
}

func TestScale(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	r.Scale(2)
	if r.Width != 20 || r.Height != 10 {
		t.Errorf("Expected 20, 10, got %d, %d", r.Width, r.Height)
	}
}
