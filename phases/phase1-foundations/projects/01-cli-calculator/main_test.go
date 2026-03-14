package main

import (
	"math"
	"testing"
)

func TestCalculate(t *testing.T) {
	tests := []struct {
		op      Operation
		a, b    float64
		want    float64
		wantErr bool
	}{
		{OpAdd, 3, 4, 7, false},
		{OpSubtract, 10, 3, 7, false},
		{OpMultiply, 4, 5, 20, false},
		{OpDivide, 10, 2, 5, false},
		{OpDivide, 5, 0, 0, true},
		{OpSqrt, 9, 0, 3, false},
		{OpSqrt, -1, 0, 0, true},
		{"unknown", 1, 2, 0, true},
	}
	for _, tt := range tests {
		result, err := calculate(tt.op, tt.a, tt.b)
		if tt.wantErr {
			if err == nil {
				t.Errorf("calculate(%s, %v, %v) expected error, got nil", tt.op, tt.a, tt.b)
			}
			continue
		}
		if err != nil {
			t.Errorf("calculate(%s, %v, %v) unexpected error: %v", tt.op, tt.a, tt.b, err)
			continue
		}
		if math.Abs(result-tt.want) > 1e-9 {
			t.Errorf("calculate(%s, %v, %v) = %v, want %v", tt.op, tt.a, tt.b, result, tt.want)
		}
	}
}

func TestParseArgs(t *testing.T) {
	_, _, _, err := parseArgs([]string{"add", "3", "4"})
	if err != nil {
		t.Errorf("parseArgs(add 3 4) unexpected error: %v", err)
	}
	_, _, _, err = parseArgs([]string{"sqrt", "9"})
	if err != nil {
		t.Errorf("parseArgs(sqrt 9) unexpected error: %v", err)
	}
	_, _, _, err = parseArgs([]string{})
	if err == nil {
		t.Error("parseArgs(empty) should return error")
	}
	_, _, _, err = parseArgs([]string{"add", "notanumber", "4"})
	if err == nil {
		t.Error("parseArgs(bad number) should return error")
	}
}
