package exercise_16

import (
	"errors"
	"testing"
)

func TestStack(t *testing.T) {
	s := &Stack{}
	if s.Len() != 0 {
		t.Errorf("new stack Len = %d, want 0", s.Len())
	}

	s.Push(1)
	s.Push("hello")
	s.Push(3.14)

	if s.Len() != 3 {
		t.Errorf("after 3 pushes, Len = %d, want 3", s.Len())
	}

	top, err := s.Peek()
	if err != nil || top != 3.14 {
		t.Errorf("Peek = (%v, %v), want (3.14, nil)", top, err)
	}

	v, err := s.Pop()
	if err != nil || v != 3.14 {
		t.Errorf("Pop = (%v, %v), want (3.14, nil)", v, err)
	}
	if s.Len() != 2 {
		t.Errorf("after Pop, Len = %d, want 2", s.Len())
	}
}

func TestStackEmpty(t *testing.T) {
	s := &Stack{}
	_, err := s.Pop()
	if !errors.Is(err, ErrEmptyStack) {
		t.Errorf("Pop empty stack should return ErrEmptyStack, got %v", err)
	}
	_, err = s.Peek()
	if !errors.Is(err, ErrEmptyStack) {
		t.Errorf("Peek empty stack should return ErrEmptyStack, got %v", err)
	}
}
