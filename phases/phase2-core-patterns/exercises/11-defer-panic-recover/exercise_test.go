package exercise_11

import (
	"errors"
	"testing"
)

func TestWithCleanup(t *testing.T) {
	cleaned := false
	err := WithCleanup(
		func() error { return nil },
		func() { cleaned = true },
	)
	if err != nil {
		t.Errorf("WithCleanup unexpected error: %v", err)
	}
	if !cleaned {
		t.Error("cleanup was not called")
	}
}

func TestWithCleanupOnError(t *testing.T) {
	cleaned := false
	myErr := errors.New("oops")
	err := WithCleanup(
		func() error { return myErr },
		func() { cleaned = true },
	)
	if !errors.Is(err, myErr) {
		t.Errorf("expected myErr, got %v", err)
	}
	if !cleaned {
		t.Error("cleanup should be called even when fn returns error")
	}
}

func TestSafeRun(t *testing.T) {
	err := SafeRun(func() {})
	if err != nil {
		t.Errorf("SafeRun(no panic) unexpected error: %v", err)
	}
	err = SafeRun(func() { panic("something went wrong") })
	if err == nil {
		t.Error("SafeRun(panic) should return an error")
	}
}

func TestDeferOrder(t *testing.T) {
	log := DeferOrder()
	if len(log) != 3 {
		t.Fatalf("DeferOrder len = %d, want 3", len(log))
	}
	// Defers execute LIFO: "third" appended first by defer, runs last...
	// actually defers run LIFO so the last defer registered runs first
	// Expected: ["third", "second", "first"]
	if log[0] != "third" || log[1] != "second" || log[2] != "first" {
		t.Errorf("DeferOrder = %v, want [third second first]", log)
	}
}
