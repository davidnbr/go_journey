package exercise_06

import (
	"sort"
	"testing"
)

func sendAll(values []int) <-chan int {
	ch := make(chan int, len(values))
	for _, v := range values {
		ch <- v
	}
	close(ch)
	return ch
}

func TestFanIn(t *testing.T) {
	ch1 := sendAll([]int{1, 2, 3})
	ch2 := sendAll([]int{4, 5, 6})
	merged := FanIn(ch1, ch2)

	var got []int
	for v := range merged {
		got = append(got, v)
	}
	sort.Ints(got)
	want := []int{1, 2, 3, 4, 5, 6}
	if len(got) != 6 {
		t.Fatalf("FanIn got %d values, want 6: %v", len(got), got)
	}
	for i, v := range got {
		if v != want[i] {
			t.Errorf("got[%d] = %d, want %d", i, v, want[i])
		}
	}
}

func TestFanOut(t *testing.T) {
	in := sendAll([]int{1, 2, 3, 4, 5, 6})
	outs := FanOut(in, 3)
	if len(outs) != 3 {
		t.Fatalf("FanOut(3) returned %d channels, want 3", len(outs))
	}
	merged := FanIn(outs...)
	var got []int
	for v := range merged {
		got = append(got, v)
	}
	if len(got) != 6 {
		t.Errorf("FanOut+FanIn got %d values, want 6", len(got))
	}
}
