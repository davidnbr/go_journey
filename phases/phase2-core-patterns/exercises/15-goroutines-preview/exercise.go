package exercise_15

import (
	"sync"
	"time"
)

// RunConcurrently runs all fns concurrently and waits for all to finish.
// TODO: use sync.WaitGroup, launch each fn as a goroutine
func RunConcurrently(fns []func()) {
	// no-op stub — runs sequentially instead
	for _, fn := range fns {
		fn()
	}
}

// IsConcurrent returns true if the goroutines ran concurrently
// (i.e., finished faster than running them sequentially would).
// This is a detection helper — don't modify.
func IsConcurrent(fns []func(sleepDur time.Duration)) bool {
	var wg sync.WaitGroup
	start := time.Now()
	for _, fn := range fns {
		wg.Add(1)
		f := fn
		go func() {
			defer wg.Done()
			f(20 * time.Millisecond)
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	// Sequential would take 5*20ms = 100ms; concurrent should be ~20ms
	return elapsed < 80*time.Millisecond
}
