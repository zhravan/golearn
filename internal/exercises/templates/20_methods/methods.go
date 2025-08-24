package methods

type Rectangle struct {
	Width, Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(factor int) {
	r.Width *= factor
	r.Height *= factor
}
