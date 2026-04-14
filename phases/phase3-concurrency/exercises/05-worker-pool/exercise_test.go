package exercise_05

import (
	"sort"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	input := []Job{
		{ID: 1, Input: 2},
		{ID: 2, Input: 3},
		{ID: 3, Input: 5},
		{ID: 4, Input: 7},
	}
	jobs := Dispatch(input)
	results := WorkerPool(3, jobs, func(n int) int { return n * n })

	var got []int
	for r := range results {
		got = append(got, r.Output)
	}
	sort.Ints(got)

	want := []int{4, 9, 25, 49}
	if len(got) != len(want) {
		t.Fatalf("WorkerPool got %d results, want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("result[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestWorkerPoolEmpty(t *testing.T) {
	jobs := Dispatch([]Job{})
	results := WorkerPool(2, jobs, func(n int) int { return n })
	count := 0
	for range results {
		count++
	}
	if count != 0 {
		t.Errorf("empty pool got %d results, want 0", count)
	}
}
