package exercise_11

import (
	"errors"
	"testing"
	"time"
)

func TestParallel(t *testing.T) {
	fns := []func() int{
		func() int { time.Sleep(10 * time.Millisecond); return 1 },
		func() int { time.Sleep(10 * time.Millisecond); return 2 },
		func() int { time.Sleep(10 * time.Millisecond); return 3 },
		func() int { time.Sleep(10 * time.Millisecond); return 4 },
		func() int { time.Sleep(10 * time.Millisecond); return 5 },
	}
	start := time.Now()
	results := Parallel(fns)
	elapsed := time.Since(start)

	for i, v := range results {
		if v != i+1 {
			t.Errorf("results[%d] = %d, want %d", i, v, i+1)
		}
	}
	if elapsed > 80*time.Millisecond {
		t.Errorf("Parallel took %v — should be concurrent (~10ms)", elapsed)
	}
}

func TestParallelWithErrors(t *testing.T) {
	myErr := errors.New("fail")
	fns := []func() error{
		func() error { return nil },
		func() error { return myErr },
		func() error { return nil },
	}
	errs := ParallelWithErrors(fns)
	if len(errs) != 3 {
		t.Fatalf("len(errs) = %d, want 3", len(errs))
	}
	if errs[0] != nil {
		t.Errorf("errs[0] = %v, want nil", errs[0])
	}
	if !errors.Is(errs[1], myErr) {
		t.Errorf("errs[1] = %v, want myErr", errs[1])
	}
	if errs[2] != nil {
		t.Errorf("errs[2] = %v, want nil", errs[2])
	}
}
