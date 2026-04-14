package exercise_04

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {
	c := Circle{Radius: 5}
	if math.Abs(c.Area()-math.Pi*25) > 1e-9 {
		t.Errorf("Circle.Area() = %v, want π*25", c.Area())
	}
	if math.Abs(c.Perimeter()-2*math.Pi*5) > 1e-9 {
		t.Errorf("Circle.Perimeter() = %v, want 2π*5", c.Perimeter())
	}
}

func TestRectangle(t *testing.T) {
	r := Rectangle{Width: 4, Height: 3}
	if r.Area() != 12 {
		t.Errorf("Rectangle.Area() = %v, want 12", r.Area())
	}
	if r.Perimeter() != 14 {
		t.Errorf("Rectangle.Perimeter() = %v, want 14", r.Perimeter())
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{
		Rectangle{Width: 2, Height: 3},
		Rectangle{Width: 4, Height: 5},
	}
	if got := TotalArea(shapes); got != 26 {
		t.Errorf("TotalArea = %v, want 26", got)
	}
}

func TestInterfaceSatisfaction(t *testing.T) {
	var _ Shape = Circle{}
	var _ Shape = Rectangle{}
}
