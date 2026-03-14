package exercise_10

import "testing"

func TestDefaultConfig(t *testing.T) {
	c := DefaultConfig()
	if c.Workers != 4 {
		t.Errorf("DefaultConfig Workers = %d, want 4", c.Workers)
	}
	if c.Host != "localhost" {
		t.Errorf("DefaultConfig Host = %q, want localhost", c.Host)
	}
}

func TestWithDebug(t *testing.T) {
	c := DefaultConfig()
	c2 := WithDebug(c)
	if !c2.Debug {
		t.Error("WithDebug should return config with Debug=true")
	}
	if c.Debug {
		t.Error("WithDebug should not modify original config")
	}
}

func TestSetWorkers(t *testing.T) {
	c := DefaultConfig()
	SetWorkers(&c, 8)
	if c.Workers != 8 {
		t.Errorf("after SetWorkers(8), Workers = %d, want 8", c.Workers)
	}
}

func TestClone(t *testing.T) {
	c := &Config{Debug: true, Workers: 2, Host: "prod"}
	cp := Clone(c)
	if cp != *c {
		t.Errorf("Clone result differs from original: %+v vs %+v", cp, *c)
	}
	cp.Workers = 99
	if c.Workers == 99 {
		t.Error("Clone should return independent copy")
	}
}
