package taskqueue

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestQueueBasic(t *testing.T) {
	q := NewQueue(3, 10)
	if q == nil {
		t.Fatal("NewQueue returned nil")
	}

	ctx := context.Background()
	q.Start(ctx)

	var completed int64
	for i := 0; i < 9; i++ {
		i := i
		err := q.Submit(Task{
			ID: i,
			Fn: func(ctx context.Context) error {
				atomic.AddInt64(&completed, 1)
				return nil
			},
		})
		if err != nil {
			t.Errorf("Submit task %d error: %v", i, err)
		}
	}

	q.Stop()

	// drain results
	for range q.Results() {
	}

	if completed != 9 {
		t.Errorf("completed = %d, want 9", completed)
	}
}

func TestQueueConcurrent(t *testing.T) {
	q := NewQueue(4, 20)
	ctx := context.Background()
	q.Start(ctx)

	start := time.Now()
	for i := 0; i < 8; i++ {
		q.Submit(Task{
			ID: i,
			Fn: func(ctx context.Context) error {
				time.Sleep(20 * time.Millisecond)
				return nil
			},
		})
	}
	q.Stop()
	for range q.Results() {
	}
	elapsed := time.Since(start)
	// 8 tasks, 4 workers, 20ms each → 2 rounds → ~40ms
	if elapsed > 200*time.Millisecond {
		t.Errorf("Queue took %v, want ~40ms with 4 workers", elapsed)
	}
}
