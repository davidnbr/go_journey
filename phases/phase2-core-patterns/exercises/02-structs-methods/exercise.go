package exercise_02

import "math"

// Circle has a center point and radius.
type Circle struct {
	X, Y, Radius float64
}

// Area returns the area of the circle: π * r²
// Use a value receiver — does not modify the circle.
// TODO: return math.Pi * c.Radius * c.Radius
func (c Circle) Area() float64 {
	return 0
}

// Scale multiplies the radius by factor.
// Use a pointer receiver — modifies the circle in-place.
// TODO: c.Radius *= factor
func (c *Circle) Scale(factor float64) {
	// no-op stub
}

// String returns a human-readable description.
// TODO: use fmt.Sprintf to return "Circle(r=<radius>)"
func (c Circle) String() string {
	return ""
}

var _ = math.Pi
