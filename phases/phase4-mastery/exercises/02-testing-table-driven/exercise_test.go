package exercise_02

import (
	"testing"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{"positive", "42", 42, false},
		{"negative", "-7", -7, false},
		{"zero", "0", 0, false},
		{"invalid", "abc", 0, true},
		{"empty", "", 0, true},
		{"float", "3.14", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseInt(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseInt(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseInt(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}

func TestSafeDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr bool
	}{
		{"positive", 10, 2, 5, false},
		{"negative", -10, 2, -5, false},
		{"zero divisor", 5, 0, 0, true},
		{"zero dividend", 0, 5, 0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SafeDivide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("SafeDivide(%d,%d) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("SafeDivide(%d,%d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
