package exercise_09

import (
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	if c.Value() != 1000 {
		t.Errorf("SafeCounter = %d, want 1000 (data race?)", c.Value())
	}
}

func TestSafeMap(t *testing.T) {
	m := NewSafeMap()
	if m == nil {
		t.Fatal("NewSafeMap returned nil")
	}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			key := "key"
			m.Set(key, i)
			m.Get(key)
		}()
	}
	wg.Wait()
}
