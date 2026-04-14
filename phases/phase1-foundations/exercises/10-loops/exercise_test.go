package exercise_10

import (
	"reflect"
	"testing"
)

func TestSumUpTo(t *testing.T) {
	tests := []struct{ n, want int }{
		{5, 15},
		{1, 1},
		{0, 0},
		{-1, 0},
		{10, 55},
	}
	for _, tt := range tests {
		got := SumUpTo(tt.n)
		if got != tt.want {
			t.Errorf("SumUpTo(%d) = %d, want %d", tt.n, got, tt.want)
		}
	}
}

func TestCollectEven(t *testing.T) {
	got := CollectEven([]int{1, 2, 3, 4, 5, 6})
	want := []int{2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("CollectEven([1..6]) = %v, want %v", got, want)
	}
}

func TestCollectEvenEmpty(t *testing.T) {
	got := CollectEven([]int{1, 3, 5})
	if len(got) != 0 {
		t.Errorf("CollectEven(all odd) should return empty, got %v", got)
	}
}

func TestFibonacci(t *testing.T) {
	got := Fibonacci(8)
	want := []int{0, 1, 1, 2, 3, 5, 8, 13}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Fibonacci(8) = %v, want %v", got, want)
	}
	if Fibonacci(0) != nil && len(Fibonacci(0)) != 0 {
		t.Error("Fibonacci(0) should return empty")
	}
}
