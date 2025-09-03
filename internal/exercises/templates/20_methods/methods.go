package methods

// TODO:
// - Add methods to a struct type.
// - Area should be a value-receiver method returning width*height.
// - Scale should be a pointer-receiver method that multiplies both fields.

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int {
	// TODO: compute area
	return 0
}

func (r *Rectangle) Scale(factor int) {
	// TODO: scale the rectangle in place
}
