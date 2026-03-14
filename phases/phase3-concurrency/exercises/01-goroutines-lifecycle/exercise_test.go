package exercise_01

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestRunAll(t *testing.T) {
	var n int64
	fns := make([]func(), 10)
	for i := range fns {
		fns[i] = func() { atomic.AddInt64(&n, 1) }
	}
	RunAll(fns)
	if n != 10 {
		t.Errorf("RunAll ran %d functions, want 10", n)
	}
}

func TestRunAllConcurrent(t *testing.T) {
	start := time.Now()
	fns := make([]func(), 5)
	for i := range fns {
		fns[i] = func() { time.Sleep(20 * time.Millisecond) }
	}
	RunAll(fns)
	elapsed := time.Since(start)
	// Sequential would be ~100ms; concurrent should be ~20ms
	if elapsed > 80*time.Millisecond {
		t.Errorf("RunAll took %v — are you running concurrently?", elapsed)
	}
}

func TestRunWithTimeout(t *testing.T) {
	// Fast fn: should complete
	ok := RunWithTimeout(func() { time.Sleep(5 * time.Millisecond) }, 100*time.Millisecond)
	if !ok {
		t.Error("RunWithTimeout: fast fn should complete in time")
	}
	// Slow fn: should time out
	ok = RunWithTimeout(func() { time.Sleep(200 * time.Millisecond) }, 20*time.Millisecond)
	if ok {
		t.Error("RunWithTimeout: slow fn should time out")
	}
}

func TestLeakFree(t *testing.T) {
	ch := make(chan int, 5)
	count := 0
	ch <- 1
	ch <- 2
	ch <- 3
	done := LeakFree(ch, &count)
	close(ch)
	select {
	case <-done:
	case <-time.After(500 * time.Millisecond):
		t.Error("LeakFree goroutine did not exit after channel closed")
	}
	if count != 3 {
		t.Errorf("count = %d, want 3", count)
	}
}
