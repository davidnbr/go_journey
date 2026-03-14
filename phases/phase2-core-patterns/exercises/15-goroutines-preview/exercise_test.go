package exercise_15

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestRunConcurrently(t *testing.T) {
	var count int64
	var mu sync.Mutex
	results := make([]int, 5)

	fns := make([]func(), 5)
	for i := range fns {
		i := i
		fns[i] = func() {
			mu.Lock()
			results[i] = i * 2
			mu.Unlock()
			atomic.AddInt64(&count, 1)
		}
	}

	RunConcurrently(fns)

	if count != 5 {
		t.Errorf("RunConcurrently ran %d functions, want 5", count)
	}
	for i, v := range results {
		if v != i*2 {
			t.Errorf("results[%d] = %d, want %d", i, v, i*2)
		}
	}
}
