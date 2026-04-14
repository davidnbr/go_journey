package exercise_02

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	c := Circle{Radius: 5}
	want := math.Pi * 25
	if got := c.Area(); math.Abs(got-want) > 1e-9 {
		t.Errorf("Area() = %v, want %v", got, want)
	}
}

func TestScale(t *testing.T) {
	c := Circle{Radius: 4}
	c.Scale(2.5)
	if math.Abs(c.Radius-10) > 1e-9 {
		t.Errorf("after Scale(2.5), Radius = %v, want 10", c.Radius)
	}
}

func TestString(t *testing.T) {
	c := Circle{Radius: 3}
	s := c.String()
	if s == "" {
		t.Error("String() should not be empty")
	}
	// Must contain the radius
	if s != "Circle(r=3)" {
		t.Errorf("String() = %q, want \"Circle(r=3)\"", s)
	}
}
