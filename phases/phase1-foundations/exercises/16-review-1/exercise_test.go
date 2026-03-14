package exercise_16

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestCelsiusToFahrenheit(t *testing.T) {
	tests := []struct{ c, want float64 }{
		{0, 32},
		{100, 212},
		{-40, -40},
	}
	for _, tt := range tests {
		got := CelsiusToFahrenheit(tt.c)
		if got != tt.want {
			t.Errorf("CelsiusToFahrenheit(%v) = %v, want %v", tt.c, got, tt.want)
		}
	}
}

func TestRateCounter(t *testing.T) {
	counter := RateCounter()
	if counter == nil {
		t.Fatal("RateCounter() returned nil")
	}
	counter("HTTP")
	counter("HTTP")
	counter("TCP")
	if got := counter("HTTP"); got != 3 {
		t.Errorf("HTTP count = %d, want 3", got)
	}
	if got := counter(""); got != 4 {
		t.Errorf("total count = %d, want 4 (3 HTTP + 1 TCP)", got)
	}
}

func TestGroupByLength(t *testing.T) {
	got := GroupByLength([]string{"go", "is", "fun", "and", "fast"})
	if len(got[2]) != 2 {
		t.Errorf("length-2 group: %v, want [go is]", got[2])
	}
	if len(got[3]) != 2 {
		t.Errorf("length-3 group: %v, want [fun and]", got[3])
	}
	if len(got[4]) != 1 {
		t.Errorf("length-4 group: %v, want [fast]", got[4])
	}
}

func TestFizzBuzz(t *testing.T) {
	got := FizzBuzz(15)
	expected := []string{
		"1", "2", "Fizz", "4", "Buzz",
		"Fizz", "7", "8", "Fizz", "Buzz",
		"11", "Fizz", "13", "14", "FizzBuzz",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("FizzBuzz(15) = %v, want %v", got, expected)
	}
}

func TestFirstWord(t *testing.T) {
	w, err := FirstWord("hello world")
	if err != nil || w != "hello" {
		t.Errorf("FirstWord(hello world) = (%q, %v), want (hello, nil)", w, err)
	}
	_, err = FirstWord("")
	if !errors.Is(err, ErrEmptyInput) {
		t.Errorf("FirstWord() should return ErrEmptyInput, got %v", err)
	}
	_, err = FirstWord("   ")
	if !errors.Is(err, ErrEmptyInput) {
		t.Errorf("FirstWord(whitespace) should return ErrEmptyInput, got %v", err)
	}
}

// ensure fmt is used
var _ = fmt.Sprintf
