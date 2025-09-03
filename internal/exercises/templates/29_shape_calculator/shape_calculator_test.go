package shape_calculator

import "testing"

func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 5}
	expectedArea := 78.53981633974483
	if c.Area() != expectedArea {
		t.Errorf("Expected area %f, got %f", expectedArea, c.Area())
	}
}

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 10, Height: 5}
	expectedArea := 50.0
	if r.Area() != expectedArea {
		t.Errorf("Expected area %f, got %f", expectedArea, r.Area())
	}
}

func TestShapeInterface(t *testing.T) {
	var s Shape

	c := Circle{Radius: 1}
	s = c
	if s.Area() == 0 {
		t.Errorf("Circle area through interface is 0")
	}

	r := Rectangle{Width: 1, Height: 1}
	s = r
	if s.Area() == 0 {
		t.Errorf("Rectangle area through interface is 0")
	}
}
