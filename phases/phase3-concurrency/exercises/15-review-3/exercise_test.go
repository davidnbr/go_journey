package exercise_15

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestSemaphore(t *testing.T) {
	s := NewSemaphore(2)
	if s == nil {
		t.Fatal("NewSemaphore returned nil")
	}
	s.Acquire()
	s.Acquire()
	// Third acquire would block — test just that acquire/release don't deadlock
	s.Release()
	s.Acquire() // re-acquire after release
	s.Release()
	s.Release()
}

func TestBoundedParallel(t *testing.T) {
	var maxConcurrent int64
	var current int64

	fns := make([]func(context.Context) error, 10)
	for i := range fns {
		fns[i] = func(ctx context.Context) error {
			n := atomic.AddInt64(&current, 1)
			for {
				prev := atomic.LoadInt64(&maxConcurrent)
				if n <= prev || atomic.CompareAndSwapInt64(&maxConcurrent, prev, n) {
					break
				}
			}
			time.Sleep(20 * time.Millisecond)
			atomic.AddInt64(&current, -1)
			return nil
		}
	}

	start := time.Now()
	errs := BoundedParallel(context.Background(), 3, fns)
	elapsed := time.Since(start)

	for i, err := range errs {
		if err != nil {
			t.Errorf("fn[%d] error: %v", i, err)
		}
	}
	// 10 fns, max 3 concurrent, 20ms each → ~4 rounds → ~80ms
	if elapsed > 300*time.Millisecond {
		t.Errorf("BoundedParallel took %v — too slow", elapsed)
	}
	if maxConcurrent > 3 {
		t.Errorf("max concurrent = %d, want <= 3", maxConcurrent)
	}
}
