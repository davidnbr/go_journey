package exercise_05

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct{ a, b, want int }{
		{1, 2, 3},
		{-1, 1, 0},
		{0, 0, 0},
	}
	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestDivide(t *testing.T) {
	q, ok := Divide(10, 2)
	if !ok || q != 5 {
		t.Errorf("Divide(10, 2) = (%d, %v), want (5, true)", q, ok)
	}
	_, ok = Divide(5, 0)
	if ok {
		t.Error("Divide(5, 0) should return ok=false")
	}
}

func TestMinMax(t *testing.T) {
	min, max := MinMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	if min != 1 {
		t.Errorf("MinMax min = %d, want 1", min)
	}
	if max != 9 {
		t.Errorf("MinMax max = %d, want 9", max)
	}
}

func TestMinMaxEmpty(t *testing.T) {
	min, max := MinMax([]int{})
	if min != 0 || max != 0 {
		t.Errorf("MinMax([]) = (%d, %d), want (0, 0)", min, max)
	}
}
