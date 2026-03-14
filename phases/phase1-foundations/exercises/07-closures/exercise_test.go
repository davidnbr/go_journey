package exercise_07

import "testing"

func TestCounter(t *testing.T) {
	c := Counter()
	if c == nil {
		t.Fatal("Counter() returned nil")
	}
	for i := 1; i <= 5; i++ {
		if got := c(); got != i {
			t.Errorf("counter call %d = %d, want %d", i, got, i)
		}
	}
}

func TestCounterIndependent(t *testing.T) {
	c1 := Counter()
	c2 := Counter()
	if c1 == nil || c2 == nil {
		t.Fatal("Counter() returned nil")
	}
	c1()
	c1()
	if got := c2(); got != 1 {
		t.Errorf("independent counter c2 first call = %d, want 1", got)
	}
}

func TestAdder(t *testing.T) {
	add5 := Adder(5)
	if add5 == nil {
		t.Fatal("Adder(5) returned nil")
	}
	if got := add5(3); got != 8 {
		t.Errorf("Adder(5)(3) = %d, want 8", got)
	}
}

func TestOnce(t *testing.T) {
	calls := 0
	fn := func() string {
		calls++
		return "result"
	}
	once := Once(fn)
	if once == nil {
		t.Fatal("Once returned nil")
	}
	r1 := once()
	r2 := once()
	r3 := once()
	if r1 != "result" || r2 != "result" || r3 != "result" {
		t.Error("Once should return the first result on all calls")
	}
	if calls != 1 {
		t.Errorf("fn was called %d times, want exactly 1", calls)
	}
}
