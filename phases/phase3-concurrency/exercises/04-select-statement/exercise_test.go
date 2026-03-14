package exercise_04

import (
	"testing"
	"time"
)

func TestFirst(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch2 <- "winner"
	got := First(ch1, ch2)
	if got != "winner" {
		t.Errorf("First = %q, want winner", got)
	}
}

func TestWithDefault(t *testing.T) {
	ch := make(chan int)
	got := WithDefault(ch, 42)
	if got != 42 {
		t.Errorf("WithDefault(empty) = %d, want 42", got)
	}
	ch2 := make(chan int, 1)
	ch2 <- 99
	got = WithDefault(ch2, 0)
	if got != 99 {
		t.Errorf("WithDefault(ready) = %d, want 99", got)
	}
}

func TestTimeout(t *testing.T) {
	ch := make(chan string, 1)
	ch <- "hello"
	val, timedOut := Timeout(ch, 100*time.Millisecond)
	if timedOut || val != "hello" {
		t.Errorf("Timeout(ready) = (%q, %v), want (hello, false)", val, timedOut)
	}
	empty := make(chan string)
	_, timedOut = Timeout(empty, 20*time.Millisecond)
	if !timedOut {
		t.Error("Timeout(empty, 20ms) should time out")
	}
}

func TestMerge(t *testing.T) {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)
	for _, v := range []int{1, 3, 5} {
		ch1 <- v
	}
	for _, v := range []int{2, 4, 6} {
		ch2 <- v
	}
	close(ch1)
	close(ch2)

	out := Merge(ch1, ch2)
	var results []int
	for v := range out {
		results = append(results, v)
	}
	if len(results) != 6 {
		t.Errorf("Merge got %d values, want 6", len(results))
	}
}
