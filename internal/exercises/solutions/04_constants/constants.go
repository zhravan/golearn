package constants

import "math"

// CircleArea calculates and returns the area of a circle with the given radius.
func CircleArea(r float64) float64 {
	return math.Pi * r * r
}
