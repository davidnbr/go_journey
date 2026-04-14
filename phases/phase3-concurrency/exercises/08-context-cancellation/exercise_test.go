package exercise_08

import (
	"context"
	"testing"
	"time"
)

func TestSleepWithContext(t *testing.T) {
	// Normal sleep completes
	ctx := context.Background()
	err := SleepWithContext(ctx, 10*time.Millisecond)
	if err != nil {
		t.Errorf("SleepWithContext(bg, 10ms) unexpected error: %v", err)
	}

	// Cancelled context returns early
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel immediately
	start := time.Now()
	err = SleepWithContext(ctx, 5*time.Second) // should not sleep 5s
	elapsed := time.Since(start)
	if err == nil {
		t.Error("SleepWithContext(cancelled) should return error")
	}
	if elapsed > 500*time.Millisecond {
		t.Errorf("SleepWithContext(cancelled) took %v — should return immediately", elapsed)
	}
}

func TestFetchAllCancelled(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	urls := make([]string, 10)
	for i := range urls {
		urls[i] = "http://example.com"
	}
	// Each fetch takes 50ms, timeout is 30ms → should get fewer than 10
	got := FetchAll(ctx, urls, 50*time.Millisecond)
	if got == len(urls) {
		t.Error("FetchAll with short timeout should not complete all fetches")
	}
}

func TestWithDeadline(t *testing.T) {
	// Fast fn: completes before deadline
	err := WithDeadline(func(ctx context.Context) error {
		return SleepWithContext(ctx, 5*time.Millisecond)
	}, 100*time.Millisecond)
	if err != nil {
		t.Errorf("WithDeadline(fast) unexpected error: %v", err)
	}

	// Slow fn: deadline exceeded
	err = WithDeadline(func(ctx context.Context) error {
		return SleepWithContext(ctx, 5*time.Second)
	}, 20*time.Millisecond)
	if err == nil {
		t.Error("WithDeadline(slow) should return deadline error")
	}
}
