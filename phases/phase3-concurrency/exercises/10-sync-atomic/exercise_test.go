package exercise_10

import (
	"sync"
	"testing"
)

func TestAtomicCounter(t *testing.T) {
	c := &AtomicCounter{}
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	if c.Value() != 10000 {
		t.Errorf("AtomicCounter = %d, want 10000", c.Value())
	}
}

func TestAtomicAdd(t *testing.T) {
	c := &AtomicCounter{}
	c.Add(5)
	c.Add(-2)
	if c.Value() != 3 {
		t.Errorf("AtomicCounter after Add(5)+Add(-2) = %d, want 3", c.Value())
	}
}

func TestCompareAndSwap(t *testing.T) {
	c := &AtomicCounter{}
	c.Add(10)
	swapped := c.CompareAndSwap(10, 20)
	if !swapped || c.Value() != 20 {
		t.Errorf("CAS(10→20) swapped=%v value=%d, want (true, 20)", swapped, c.Value())
	}
	swapped = c.CompareAndSwap(10, 99) // should fail — value is 20
	if swapped {
		t.Error("CAS(10→99) should fail when value is 20")
	}
}
