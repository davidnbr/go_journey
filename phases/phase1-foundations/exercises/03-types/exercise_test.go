package exercise_03

import "testing"

func TestIntToFloat(t *testing.T) {
	got := IntToFloat(5)
	want := float64(5)
	if got != want {
		t.Errorf("IntToFloat(5) = %v, want %v", got, want)
	}
}

func TestFloatToInt(t *testing.T) {
	tests := []struct {
		input float64
		want  int
	}{
		{3.7, 3},
		{-1.9, -1},
		{0.0, 0},
	}
	for _, tt := range tests {
		got := FloatToInt(tt.input)
		if got != tt.want {
			t.Errorf("FloatToInt(%v) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestStringLength(t *testing.T) {
	if got := StringLength("hello"); got != 5 {
		t.Errorf("StringLength(\"hello\") = %d, want 5", got)
	}
	if got := StringLength(""); got != 0 {
		t.Errorf("StringLength(\"\") = %d, want 0", got)
	}
}

func TestIsPositive(t *testing.T) {
	if !IsPositive(1) {
		t.Error("IsPositive(1) should be true")
	}
	if IsPositive(0) {
		t.Error("IsPositive(0) should be false")
	}
	if IsPositive(-5) {
		t.Error("IsPositive(-5) should be false")
	}
}
