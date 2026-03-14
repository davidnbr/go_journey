package exercise_18

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

// EventBus dispatches events to registered handlers.
type EventBus struct {
	mu       sync.RWMutex
	handlers map[string][]func(data interface{})
}

// NewEventBus creates a new EventBus.
// TODO: initialize handlers map
func NewEventBus() *EventBus {
	return &EventBus{}
}

// Subscribe registers a handler for an event type.
// TODO: append handler to handlers[event]
func (b *EventBus) Subscribe(event string, handler func(data interface{})) {
	// no-op
}

// Publish sends data to all handlers registered for event.
// TODO: RLock, get handlers, RUnlock, call each handler
func (b *EventBus) Publish(event string, data interface{}) {
	// no-op
}

// ErrWriterFull is returned when the writer is full.
var ErrWriterFull = errors.New("writer full")

// LimitedWriter writes to w but returns ErrWriterFull after limit bytes.
type LimitedWriter struct {
	W       io.Writer
	Limit   int
	written int
}

// Write implements io.Writer. Returns ErrWriterFull when limit exceeded.
// TODO: track written bytes, return ErrWriterFull when limit reached
func (lw *LimitedWriter) Write(p []byte) (int, error) {
	return 0, fmt.Errorf("not implemented")
}
