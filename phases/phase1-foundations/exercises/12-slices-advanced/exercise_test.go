package exercise_12

import (
	"reflect"
	"testing"
)

func TestDeduplicate(t *testing.T) {
	got := Deduplicate([]int{1, 2, 2, 3, 1, 4})
	want := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Deduplicate = %v, want %v", got, want)
	}
	if len(Deduplicate([]int{})) != 0 {
		t.Error("Deduplicate(empty) should be empty")
	}
}

func TestFlatten(t *testing.T) {
	got := Flatten([][]int{{1, 2}, {3, 4}, {5}})
	want := []int{1, 2, 3, 4, 5}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Flatten = %v, want %v", got, want)
	}
}

func TestCopySlice(t *testing.T) {
	orig := []int{1, 2, 3}
	cp := CopySlice(orig)
	if !reflect.DeepEqual(orig, cp) {
		t.Errorf("CopySlice content mismatch: %v vs %v", orig, cp)
	}
	cp[0] = 99
	if orig[0] == 99 {
		t.Error("CopySlice returned a view, not a copy — modifying it changed the original")
	}
}
