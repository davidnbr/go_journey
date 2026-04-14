package exercise_13

import (
	"reflect"
	"testing"
)

func TestWordCount(t *testing.T) {
	got := WordCount([]string{"go", "is", "go", "fun", "go"})
	want := map[string]int{"go": 3, "is": 1, "fun": 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("WordCount = %v, want %v", got, want)
	}
}

func TestWordCountEmpty(t *testing.T) {
	got := WordCount([]string{})
	if got == nil {
		t.Error("WordCount(empty) should return empty map, not nil")
	}
	if len(got) != 0 {
		t.Errorf("WordCount(empty) = %v, want empty map", got)
	}
}

func TestLookup(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	v, ok := Lookup(m, "a")
	if !ok || v != 1 {
		t.Errorf("Lookup(a) = (%d, %v), want (1, true)", v, ok)
	}
	_, ok = Lookup(m, "missing")
	if ok {
		t.Error("Lookup(missing) ok should be false")
	}
}

func TestInvert(t *testing.T) {
	m := map[string]string{"a": "1", "b": "2"}
	got := Invert(m)
	want := map[string]string{"1": "a", "2": "b"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Invert = %v, want %v", got, want)
	}
}
