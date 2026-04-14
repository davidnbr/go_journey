package exercise_03

import (
	"strings"
	"testing"
)

func TestConcatLoop(t *testing.T) {
	got := ConcatLoop("ab", 3)
	if got != "ababab" {
		t.Errorf("ConcatLoop = %q, want ababab", got)
	}
}

func TestConcatBuilder(t *testing.T) {
	got := ConcatBuilder("ab", 3)
	if got != "ababab" {
		t.Errorf("ConcatBuilder = %q, want ababab", got)
	}
}

func TestConcatRepeat(t *testing.T) {
	got := ConcatRepeat("ab", 3)
	if got != "ababab" {
		t.Errorf("ConcatRepeat = %q, want ababab", got)
	}
}

// Benchmarks — run with: go test -bench=. ./phases/phase4-mastery/exercises/03-testing-benchmarks/

func BenchmarkConcatLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatLoop("x", 1000)
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatBuilder("x", 1000)
	}
}

func BenchmarkConcatRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatRepeat("x", 1000)
	}
}

var _ = strings.Builder{}
