package exercise_08

import (
	"context"
	"time"
)

// SleepWithContext sleeps for d but returns early if ctx is cancelled.
// Returns ctx.Err() if cancelled, nil if sleep completed.
// TODO: select time.After(d) vs ctx.Done()
func SleepWithContext(ctx context.Context, d time.Duration) error {
	time.Sleep(d) // stub: ignores context
	return nil
}

// FetchAll simulates fetching N URLs concurrently.
// Each "fetch" takes fetchDur. If ctx is cancelled before all finish,
// remaining fetches are skipped.
// Returns the number of successful fetches.
// TODO: goroutine per URL, select ctx.Done() vs fetch completion
func FetchAll(ctx context.Context, urls []string, fetchDur time.Duration) int {
	return len(urls) // stub: always succeeds
}

// WithDeadline wraps fn in a context with the given timeout.
// Returns the result of fn and any context error.
// TODO: context.WithTimeout, defer cancel, call fn(ctx)
func WithDeadline(fn func(ctx context.Context) error, timeout time.Duration) error {
	return fn(context.Background()) // stub: no timeout
}

var _ = context.WithCancel
