package exercise_06

import (
	"testing"
)

type myStringer struct{ val string }

func (m myStringer) String() string { return m.val }

func TestStringify(t *testing.T) {
	tests := []struct {
		v    interface{}
		want string
	}{
		{"hello", "hello"},
		{42, "42"},
		{true, "true"},
		{false, "false"},
		{myStringer{"custom"}, "custom"},
	}
	for _, tt := range tests {
		got := Stringify(tt.v)
		if got != tt.want {
			t.Errorf("Stringify(%v) = %q, want %q", tt.v, got, tt.want)
		}
	}
}

func TestMustString(t *testing.T) {
	got := MustString("hello")
	if got != "hello" {
		t.Errorf("MustString(\"hello\") = %q, want \"hello\"", got)
	}
}

func TestMustStringPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("MustString(42) should panic")
		}
	}()
	MustString(42)
}

func TestAsInt(t *testing.T) {
	v, ok := AsInt(5)
	if !ok || v != 5 {
		t.Errorf("AsInt(5) = (%d, %v), want (5, true)", v, ok)
	}
	_, ok = AsInt("nope")
	if ok {
		t.Error("AsInt(string) should return ok=false")
	}
}
