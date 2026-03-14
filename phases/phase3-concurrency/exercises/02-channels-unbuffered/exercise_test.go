package exercise_02

import (
	"reflect"
	"testing"
)

func TestSend(t *testing.T) {
	ch := Send(5)
	got := Collect(ch)
	want := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Send(5) = %v, want %v", got, want)
	}
}

func TestDouble(t *testing.T) {
	in := Send(4)
	out := Double(in)
	got := Collect(out)
	want := []int{0, 2, 4, 6}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Double(0..3) = %v, want %v", got, want)
	}
}

func TestCollectEmpty(t *testing.T) {
	ch := make(chan int)
	close(ch)
	got := Collect(ch)
	if len(got) != 0 {
		t.Errorf("Collect(closed empty) = %v, want []", got)
	}
}
