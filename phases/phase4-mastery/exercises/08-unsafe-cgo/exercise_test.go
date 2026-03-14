package exercise_08

import (
	"testing"
)

func TestSizeOf(t *testing.T) {
	// int on 64-bit = 8 bytes
	size := SizeOf(int(0))
	if size == 0 {
		t.Error("SizeOf(int) should not be 0")
	}
	// bool = 1 byte
	if got := SizeOf(bool(false)); got != 1 {
		t.Errorf("SizeOf(bool) = %d, want 1", got)
	}
}

func TestAlignOf(t *testing.T) {
	a := AlignOf(int64(0))
	if a == 0 {
		t.Error("AlignOf should not be 0")
	}
}

func TestStringToBytes(t *testing.T) {
	s := "hello"
	b := StringToBytes(s)
	if string(b) != s {
		t.Errorf("StringToBytes(%q) = %q, want %q", s, string(b), s)
	}
	if len(b) != len(s) {
		t.Errorf("StringToBytes len = %d, want %d", len(b), len(s))
	}
}
