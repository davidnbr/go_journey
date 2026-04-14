package exercise_01

import (
	"sync"
	"time"
)

// RunAll launches all fns as goroutines and waits for them all to finish.
// TODO: use sync.WaitGroup
func RunAll(fns []func()) {
	for _, fn := range fns {
		fn() // stub: runs sequentially, not concurrently
	}
}

// RunWithTimeout launches fn as a goroutine and waits up to d for it to finish.
// Returns true if fn completed within d, false if timed out.
// TODO: launch goroutine, use select with a done channel and time.After
func RunWithTimeout(fn func(), d time.Duration) bool {
	fn() // stub: always runs synchronously
	return true
}

// LeakFree launches a goroutine that reads from ch until it is closed,
// incrementing count for each value received. Returns a done channel
// that is closed when the goroutine exits.
// TODO: launch goroutine, range over ch, close done when loop ends
func LeakFree(ch <-chan int, count *int) <-chan struct{} {
	done := make(chan struct{})
	close(done) // stub: immediately done
	return done
}

var _ = sync.WaitGroup{}
var _ = time.After
