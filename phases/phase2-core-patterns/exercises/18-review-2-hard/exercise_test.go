package exercise_18

import (
	"bytes"
	"errors"
	"sync"
	"testing"
)

func TestEventBus(t *testing.T) {
	bus := NewEventBus()
	if bus == nil {
		t.Fatal("NewEventBus returned nil")
	}

	var mu sync.Mutex
	received := []interface{}{}

	bus.Subscribe("deploy", func(data interface{}) {
		mu.Lock()
		received = append(received, data)
		mu.Unlock()
	})
	bus.Subscribe("deploy", func(data interface{}) {
		mu.Lock()
		received = append(received, "handler2")
		mu.Unlock()
	})

	bus.Publish("deploy", "v1.0.0")
	bus.Publish("other", "ignored")

	mu.Lock()
	defer mu.Unlock()
	if len(received) != 2 {
		t.Errorf("received %d events, want 2", len(received))
	}
}

func TestLimitedWriter(t *testing.T) {
	var buf bytes.Buffer
	lw := &LimitedWriter{W: &buf, Limit: 5}

	n, err := lw.Write([]byte("hello"))
	if err != nil || n != 5 {
		t.Errorf("Write(hello) = (%d, %v), want (5, nil)", n, err)
	}

	_, err = lw.Write([]byte("x"))
	if !errors.Is(err, ErrWriterFull) {
		t.Errorf("Write after limit should return ErrWriterFull, got %v", err)
	}
}
