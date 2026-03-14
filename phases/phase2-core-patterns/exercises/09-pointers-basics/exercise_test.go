package exercise_09

import "testing"

func TestIncrement(t *testing.T) {
	n := 5
	Increment(&n)
	if n != 6 {
		t.Errorf("after Increment, n = %d, want 6", n)
	}
}

func TestSwap(t *testing.T) {
	a, b := 10, 20
	Swap(&a, &b)
	if a != 20 || b != 10 {
		t.Errorf("after Swap: a=%d, b=%d, want a=20, b=10", a, b)
	}
}

func TestNewInt(t *testing.T) {
	p := NewInt(42)
	if p == nil {
		t.Fatal("NewInt returned nil")
	}
	if *p != 42 {
		t.Errorf("*NewInt(42) = %d, want 42", *p)
	}
}

func TestSafeDeref(t *testing.T) {
	n := 7
	if got := SafeDeref(&n, 0); got != 7 {
		t.Errorf("SafeDeref(&7) = %d, want 7", got)
	}
	if got := SafeDeref(nil, 99); got != 99 {
		t.Errorf("SafeDeref(nil, 99) = %d, want 99", got)
	}
}
