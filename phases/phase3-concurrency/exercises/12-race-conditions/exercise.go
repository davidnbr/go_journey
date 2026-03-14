package exercise_12

import "sync"

// UnsafeCounter has a data race — do NOT fix it (it's intentional for the test).
// The test for this exercises uses -race to demonstrate the problem,
// then tests SafeCounter which IS fixed.
type UnsafeCounter struct {
	count int // no protection
}

func (c *UnsafeCounter) Inc()       { c.count++ }
func (c *UnsafeCounter) Value() int { return c.count }

// SafeCounter fixes the race using a mutex.
// TODO: add sync.Mutex, lock in Inc and Value
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc increments safely.
// TODO: Lock/Unlock around increment
func (c *SafeCounter) Inc() {
	c.count++ // stub: still has race
}

// Value reads safely.
// TODO: Lock/Unlock around read
func (c *SafeCounter) Value() int {
	return c.count // stub: still has race
}

// RaceFreeSumWithChannels sums all values sent on ch using only channels (no mutex).
// Returns the sum via a result channel.
// TODO: single goroutine reads from ch, accumulates sum, sends result
func RaceFreeSumWithChannels(ch <-chan int) <-chan int {
	result := make(chan int, 1)
	result <- 0 // stub
	return result
}
