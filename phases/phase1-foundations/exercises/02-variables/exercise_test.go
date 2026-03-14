package exercise_02

import "testing"

func TestZeroValues(t *testing.T) {
	i, f, b, s := ZeroValues()
	if i != 0 {
		t.Errorf("int zero value: got %d, want 0", i)
	}
	if f != 0.0 {
		t.Errorf("float64 zero value: got %f, want 0.0", f)
	}
	if b != false {
		t.Errorf("bool zero value: got %v, want false", b)
	}
	if s != "" {
		t.Errorf("string zero value: got %q, want \"\"", s)
	}
}

func TestSwap(t *testing.T) {
	a, b := Swap("hello", "world")
	if a != "world" || b != "hello" {
		t.Errorf("Swap(\"hello\", \"world\") = (%q, %q), want (\"world\", \"hello\")", a, b)
	}
}

func TestShortDeclaration(t *testing.T) {
	got := ShortDeclaration()
	if got != 30 {
		t.Errorf("ShortDeclaration() = %d, want 30", got)
	}
}
