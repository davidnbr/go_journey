package exercise_03

import (
	"reflect"
	"testing"
	"time"
)

func TestBuffer(t *testing.T) {
	ch := Buffer([]int{1, 2, 3}, 5)
	if ch == nil {
		t.Fatal("Buffer returned nil")
	}
	if len(ch) != 3 {
		t.Errorf("buffered channel len = %d, want 3", len(ch))
	}
	if cap(ch) != 5 {
		t.Errorf("channel cap = %d, want 5", cap(ch))
	}
}

func TestBatch(t *testing.T) {
	ch := make(chan int, 7)
	for i := 1; i <= 7; i++ {
		ch <- i
	}
	close(ch)
	batches := Batch(ch, 3)
	if len(batches) != 3 {
		t.Fatalf("Batch len = %d, want 3 batches", len(batches))
	}
	if !reflect.DeepEqual(batches[0], []int{1, 2, 3}) {
		t.Errorf("batch[0] = %v, want [1 2 3]", batches[0])
	}
	if !reflect.DeepEqual(batches[2], []int{7}) {
		t.Errorf("batch[2] = %v, want [7]", batches[2])
	}
}

func TestRateLimit(t *testing.T) {
	start := time.Now()
	ch := RateLimit(3)
	// drain all 3 tokens — we just check they arrive
	count := 0
	for range ch {
		count++
	}
	_ = time.Since(start)
	if count != 3 {
		t.Errorf("RateLimit sent %d tokens, want 3", count)
	}
}
