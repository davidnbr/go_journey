package exercise_08

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		n    int
		want string
	}{
		{5, "positive"},
		{0, "zero"},
		{-3, "negative"},
	}
	for _, tt := range tests {
		got := Classify(tt.n)
		if got != tt.want {
			t.Errorf("Classify(%d) = %q, want %q", tt.n, got, tt.want)
		}
	}
}

func TestAbsoluteValue(t *testing.T) {
	tests := []struct{ n, want int }{
		{5, 5},
		{-5, 5},
		{0, 0},
	}
	for _, tt := range tests {
		got := AbsoluteValue(tt.n)
		if got != tt.want {
			t.Errorf("AbsoluteValue(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestParseScore(t *testing.T) {
	tests := []struct {
		score int
		want  string
	}{
		{95, "A"},
		{85, "B"},
		{75, "C"},
		{65, "D"},
		{55, "F"},
		{101, "invalid"},
		{-1, "invalid"},
		{100, "A"},
		{0, "F"},
	}
	for _, tt := range tests {
		got := ParseScore(tt.score)
		if got != tt.want {
			t.Errorf("ParseScore(%d) = %q, want %q", tt.score, got, tt.want)
		}
	}
}
