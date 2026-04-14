package exercise_09

import "sync"

// SafeCounter is a goroutine-safe counter.
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

// Inc increments the counter.
// TODO: Lock, increment, Unlock
func (c *SafeCounter) Inc() {
	c.count++ // stub: not thread-safe
}

// Value returns the current count.
// TODO: Lock, read, Unlock
func (c *SafeCounter) Value() int {
	return c.count // stub: not thread-safe
}

// SafeMap is a goroutine-safe map[string]int.
type SafeMap struct {
	mu sync.RWMutex
	m  map[string]int
}

// NewSafeMap creates an initialized SafeMap.
// TODO: initialize the inner map
func NewSafeMap() *SafeMap {
	return &SafeMap{} // stub: nil map will panic on write
}

// Set writes key=value under a write lock.
// TODO: Lock, set, Unlock
func (s *SafeMap) Set(key string, value int) {
	s.m[key] = value // stub: not locked
}

// Get reads key under a read lock.
// TODO: RLock, read, RUnlock
func (s *SafeMap) Get(key string) (int, bool) {
	v, ok := s.m[key]
	return v, ok // stub: not locked
}
