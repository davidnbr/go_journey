package exercise_01

import "testing"

func TestNewPoint(t *testing.T) {
	p := NewPoint(3, 4)
	if p.X != 3 || p.Y != 4 {
		t.Errorf("NewPoint(3,4) = %+v, want {3 4}", p)
	}
}

func TestArea(t *testing.T) {
	r := Rectangle{TopLeft: Point{0, 0}, Width: 5, Height: 3}
	if got := r.Area(); got != 15 {
		t.Errorf("Area() = %v, want 15", got)
	}
}

func TestContains(t *testing.T) {
	r := Rectangle{TopLeft: Point{0, 0}, Width: 10, Height: 10}
	if !r.Contains(Point{5, 5}) {
		t.Error("Contains(5,5) should be true")
	}
	if !r.Contains(Point{0, 0}) {
		t.Error("Contains(0,0) boundary should be true")
	}
	if r.Contains(Point{11, 5}) {
		t.Error("Contains(11,5) should be false")
	}
}
