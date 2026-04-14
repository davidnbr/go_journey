package exercise_16

import (
	"errors"
	"fmt"
)

// Stack is a generic LIFO data structure backed by a slice.
// TODO: add Items field []interface{} or use interface{}
type Stack struct {
	items []interface{}
}

// Push adds an item to the top of the stack.
// TODO: append to items
func (s *Stack) Push(item interface{}) {
	// no-op
}

// Pop removes and returns the top item.
// Returns ErrEmptyStack if the stack is empty.
// TODO: check len, return last item, shrink slice
func (s *Stack) Pop() (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

// Peek returns the top item without removing it.
// Returns ErrEmptyStack if empty.
func (s *Stack) Peek() (interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

// Len returns the number of items.
func (s *Stack) Len() int {
	return 0
}

// ErrEmptyStack is returned when operating on an empty stack.
var ErrEmptyStack = errors.New("stack is empty")
