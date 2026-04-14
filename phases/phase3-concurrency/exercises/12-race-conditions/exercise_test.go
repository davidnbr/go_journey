package exercise_12

import (
	"sync"
	"testing"
)

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() { defer wg.Done(); c.Inc() }()
	}
	wg.Wait()
	if c.Value() != 500 {
		t.Errorf("SafeCounter = %d, want 500", c.Value())
	}
}

func TestRaceFreeSumWithChannels(t *testing.T) {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	close(ch)

	resultCh := RaceFreeSumWithChannels(ch)
	got := <-resultCh
	if got != 15 {
		t.Errorf("RaceFreeSumWithChannels = %d, want 15", got)
	}
}
