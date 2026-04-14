package exercise_01

import "testing"

func TestIsPalindrome(t *testing.T) {
	// IsPalindrome is already implemented — test it to understand the pattern
	tests := []struct {
		input string
		want  bool
	}{
		{"racecar", true},
		{"hello", false},
		{"a man a plan a canal panama", true},
		{"", true},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := IsPalindrome(tt.input)
			if got != tt.want {
				t.Errorf("IsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct{ n, want int }{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{7, 13},
		{10, 55},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := Fibonacci(tt.n)
			if got != tt.want {
				t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestClamp(t *testing.T) {
	if got := Clamp(5, 0, 10); got != 5 {
		t.Errorf("Clamp(5,0,10) = %d, want 5", got)
	}
	if got := Clamp(-1, 0, 10); got != 0 {
		t.Errorf("Clamp(-1,0,10) = %d, want 0", got)
	}
	if got := Clamp(15, 0, 10); got != 10 {
		t.Errorf("Clamp(15,0,10) = %d, want 10", got)
	}
}
