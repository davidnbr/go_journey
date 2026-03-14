package exercise_15

import (
	"context"
	"sync"
)

// Semaphore limits concurrent access to n slots.
type Semaphore struct {
	ch chan struct{}
}

// NewSemaphore creates a semaphore with n slots.
// TODO: make(chan struct{}, n)
func NewSemaphore(n int) *Semaphore {
	return &Semaphore{}
}

// Acquire takes a slot (blocks if all slots are taken).
// TODO: send to ch
func (s *Semaphore) Acquire() {
	// no-op stub
}

// Release frees a slot.
// TODO: receive from ch
func (s *Semaphore) Release() {
	// no-op stub
}

// BoundedParallel runs all fns with at most maxConcurrent running simultaneously.
// Uses Semaphore internally.
// TODO: WaitGroup + Semaphore.Acquire/Release around each fn
func BoundedParallel(ctx context.Context, maxConcurrent int, fns []func(context.Context) error) []error {
	errs := make([]error, len(fns))
	for i, fn := range fns {
		errs[i] = fn(ctx) // stub: sequential
	}
	return errs
}

var _ = sync.WaitGroup{}
