package exercise_06

import "testing"

func TestSum(t *testing.T) {
	if got := Sum(1, 2, 3); got != 6 {
		t.Errorf("Sum(1,2,3) = %d, want 6", got)
	}
	if got := Sum(); got != 0 {
		t.Errorf("Sum() = %d, want 0", got)
	}
	if got := Sum(5); got != 5 {
		t.Errorf("Sum(5) = %d, want 5", got)
	}
}

func TestApply(t *testing.T) {
	double := func(n int) int { return n * 2 }
	got := Apply(5, double)
	if got != 10 {
		t.Errorf("Apply(5, double) = %d, want 10", got)
	}
}

func TestMultiplier(t *testing.T) {
	triple := Multiplier(3)
	if triple == nil {
		t.Fatal("Multiplier(3) returned nil")
	}
	if got := triple(4); got != 12 {
		t.Errorf("Multiplier(3)(4) = %d, want 12", got)
	}
	if got := triple(0); got != 0 {
		t.Errorf("Multiplier(3)(0) = %d, want 0", got)
	}
}
