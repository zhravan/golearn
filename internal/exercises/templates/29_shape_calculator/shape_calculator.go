package shape_calculator

import "math"

// TODO:
// - Define a Shape interface with an Area() float64 method.
// - Implement Circle and Rectangle types that satisfy Shape.
// - Implement Area for each shape using correct formulas.

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	// TODO: return circle area
	_ = math.Pi
	return 0
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	// TODO: return rectangle area
	return 0
}
