package exercise_15

import (
	"errors"
	"math"
	"testing"
)

func TestErrDivisionByZero(t *testing.T) {
	if ErrDivisionByZero == nil {
		t.Fatal("ErrDivisionByZero must not be nil")
	}
	if ErrDivisionByZero.Error() == "" {
		t.Error("ErrDivisionByZero must have a non-empty message")
	}
}

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("Divide(10,2) unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Divide(10,2) = %v, want 5.0", result)
	}

	_, err = Divide(5, 0)
	if err == nil {
		t.Fatal("Divide(5,0) should return an error")
	}
	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("Divide(5,0) error should be ErrDivisionByZero, got: %v", err)
	}
}

func TestParsePositive(t *testing.T) {
	v, err := ParsePositive(5)
	if err != nil || v != 5 {
		t.Errorf("ParsePositive(5) = (%d, %v), want (5, nil)", v, err)
	}
	_, err = ParsePositive(0)
	if err == nil {
		t.Error("ParsePositive(0) should return error")
	}
	_, err = ParsePositive(-3)
	if err == nil {
		t.Error("ParsePositive(-3) should return error")
	}
}

func TestSafeSqrt(t *testing.T) {
	result, err := SafeSqrt(4.0)
	if err != nil {
		t.Fatalf("SafeSqrt(4) unexpected error: %v", err)
	}
	if math.Abs(result-2.0) > 1e-9 {
		t.Errorf("SafeSqrt(4) = %v, want 2.0", result)
	}
	_, err = SafeSqrt(-1)
	if err == nil {
		t.Error("SafeSqrt(-1) should return error")
	}
}
