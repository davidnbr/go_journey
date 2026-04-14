package exercise_11

import (
	"reflect"
	"testing"
)

func TestFirstThree(t *testing.T) {
	got := FirstThree([]int{1, 2, 3, 4, 5})
	want := []int{1, 2, 3}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FirstThree([1..5]) = %v, want %v", got, want)
	}
	short := FirstThree([]int{1, 2})
	if !reflect.DeepEqual(short, []int{1, 2}) {
		t.Errorf("FirstThree([1,2]) = %v, want [1 2]", short)
	}
}

func TestArraySum(t *testing.T) {
	got := ArraySum([5]int{1, 2, 3, 4, 5})
	if got != 15 {
		t.Errorf("ArraySum = %d, want 15", got)
	}
}

func TestSliceInfo(t *testing.T) {
	s := make([]int, 3, 10)
	l, c := SliceInfo(s)
	if l != 3 {
		t.Errorf("length = %d, want 3", l)
	}
	if c != 10 {
		t.Errorf("capacity = %d, want 10", c)
	}
}

func TestRemoveLast(t *testing.T) {
	got := RemoveLast([]int{1, 2, 3})
	want := []int{1, 2}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("RemoveLast([1,2,3]) = %v, want %v", got, want)
	}
	empty := RemoveLast([]int{})
	if len(empty) != 0 {
		t.Errorf("RemoveLast([]) should return empty, got %v", empty)
	}
}
