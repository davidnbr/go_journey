package exercise_01

// Point represents a 2D point.
type Point struct {
	X, Y float64
}

// Rectangle has a top-left point plus width and height.
type Rectangle struct {
	TopLeft Point
	Width   float64
	Height  float64
}

// NewPoint constructs a Point.
// TODO: return Point{X: x, Y: y}
func NewPoint(x, y float64) Point {
	return Point{}
}

// Area returns the area of a Rectangle.
// TODO: return Width * Height
func (r Rectangle) Area() float64 {
	return 0
}

// Contains reports whether pt is inside (or on the boundary of) r.
// TODO: check if pt.X in [r.TopLeft.X, r.TopLeft.X+r.Width] and pt.Y similarly
func (r Rectangle) Contains(pt Point) bool {
	return false
}
