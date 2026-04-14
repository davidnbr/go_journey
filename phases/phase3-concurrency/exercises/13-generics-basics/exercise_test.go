package exercise_13

import (
	"reflect"
	"testing"
)

func TestMap(t *testing.T) {
	doubled := Map([]int{1, 2, 3}, func(n int) int { return n * 2 })
	if !reflect.DeepEqual(doubled, []int{2, 4, 6}) {
		t.Errorf("Map double = %v, want [2 4 6]", doubled)
	}
	strs := Map([]int{1, 2, 3}, func(n int) string { return "x" })
	if len(strs) != 3 {
		t.Errorf("Map to string len = %d, want 3", len(strs))
	}
}

func TestFilter(t *testing.T) {
	got := Filter([]int{1, 2, 3, 4, 5}, func(n int) bool { return n%2 == 0 })
	if !reflect.DeepEqual(got, []int{2, 4}) {
		t.Errorf("Filter evens = %v, want [2 4]", got)
	}
	strs := Filter([]string{"go", "is", "great"}, func(s string) bool { return len(s) > 2 })
	if !reflect.DeepEqual(strs, []string{"great"}) {
		t.Errorf("Filter long strings = %v, want [great]", strs)
	}
}

func TestReduce(t *testing.T) {
	sum := Reduce([]int{1, 2, 3, 4}, 0, func(acc, v int) int { return acc + v })
	if sum != 10 {
		t.Errorf("Reduce sum = %d, want 10", sum)
	}
}

func TestContains(t *testing.T) {
	if !Contains([]string{"a", "b", "c"}, "b") {
		t.Error("Contains(b) should be true")
	}
	if Contains([]int{1, 2, 3}, 4) {
		t.Error("Contains(4) should be false")
	}
}
