package exercise_10

import "sync/atomic"

// AtomicCounter uses sync/atomic for lock-free increment.
type AtomicCounter struct {
	value int64
}

// Inc atomically increments by 1.
// TODO: atomic.AddInt64
func (c *AtomicCounter) Inc() {
	c.value++ // stub: not atomic
}

// Add atomically adds delta.
// TODO: atomic.AddInt64
func (c *AtomicCounter) Add(delta int64) {
	c.value += delta // stub
}

// Value atomically loads the current value.
// TODO: atomic.LoadInt64
func (c *AtomicCounter) Value() int64 {
	return c.value // stub
}

// CompareAndSwap atomically sets value to newVal if it equals oldVal.
// Returns true if the swap happened.
// TODO: atomic.CompareAndSwapInt64
func (c *AtomicCounter) CompareAndSwap(oldVal, newVal int64) bool {
	return false // stub
}

var _ = atomic.AddInt64
