package exercise_09

import (
	"sort"
	"testing"
)

func TestSumSquares(t *testing.T) {
	// 1^2 + 2^2 + 3^2 = 1+4+9 = 14
	if got := SumSquares(3); got != 14 {
		t.Errorf("SumSquares(3) = %d, want 14", got)
	}
	if got := SumSquares(10); got != 385 {
		t.Errorf("SumSquares(10) = %d, want 385", got)
	}
}

func TestSumSquaresFast(t *testing.T) {
	for n := 1; n <= 100; n++ {
		naive := SumSquares(n)
		fast := SumSquaresFast(n)
		if naive != fast {
			t.Errorf("SumSquares(%d)=%d != SumSquaresFast(%d)=%d", n, naive, n, fast)
		}
	}
}

func TestContainsLinear(t *testing.T) {
	s := []int{3, 1, 4, 1, 5, 9, 2, 6}
	if !ContainsLinear(s, 5) {
		t.Error("ContainsLinear(5) should be true")
	}
	if ContainsLinear(s, 7) {
		t.Error("ContainsLinear(7) should be false")
	}
}

func TestContainsBinary(t *testing.T) {
	s := []int{1, 2, 3, 5, 8, 13, 21}
	if !ContainsBinary(s, 8) {
		t.Error("ContainsBinary(8) should be true")
	}
	if ContainsBinary(s, 4) {
		t.Error("ContainsBinary(4) should be false")
	}
}

func BenchmarkSumSquaresNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSquares(10000)
	}
}

func BenchmarkSumSquaresFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumSquaresFast(10000)
	}
}

func BenchmarkContainsLinear(b *testing.B) {
	s := make([]int, 10000)
	for i := range s {
		s[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsLinear(s, 9999)
	}
}

func BenchmarkContainsBinary(b *testing.B) {
	s := make([]int, 10000)
	for i := range s {
		s[i] = i
	}
	sort.Ints(s)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ContainsBinary(s, 9999)
	}
}

var _ = sort.SearchInts
