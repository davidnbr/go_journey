package exercise_04

import "math"

// Shape is the interface all shapes implement.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle implements Shape.
type Circle struct {
	Radius float64
}

// TODO: implement Area() and Perimeter() on Circle
// Area = π*r², Perimeter = 2*π*r
func (c Circle) Area() float64      { return 0 }
func (c Circle) Perimeter() float64 { return 0 }

// Rectangle implements Shape.
type Rectangle struct {
	Width, Height float64
}

// TODO: implement Area() and Perimeter() on Rectangle
// Area = W*H, Perimeter = 2*(W+H)
func (r Rectangle) Area() float64      { return 0 }
func (r Rectangle) Perimeter() float64 { return 0 }

// TotalArea returns the sum of areas of all shapes.
// TODO: iterate and sum s.Area()
func TotalArea(shapes []Shape) float64 {
	return 0
}

var _ = math.Pi
