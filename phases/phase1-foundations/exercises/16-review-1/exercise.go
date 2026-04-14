package exercise_16

import (
	"errors"
	"strings"
)

// --- Variables & Types ---

// CelsiusToFahrenheit converts Celsius to Fahrenheit: F = C*9/5 + 32
// TODO: implement the formula using float64 arithmetic
func CelsiusToFahrenheit(c float64) float64 {
	return 0
}

// --- Functions & Closures ---

// RateCounter returns a counter function that tracks counts by string category.
// Calling the returned function with a category increments its count.
// Calling with "" returns the total count across all categories.
// TODO: use a closure over a map[string]int
func RateCounter() func(category string) int {
	return nil
}

// --- Slices & Maps ---

// GroupByLength groups strings by their length.
// Returns a map[int][]string where the key is string length.
// TODO: build map, append each word to the correct length bucket
func GroupByLength(words []string) map[int][]string {
	return nil
}

// --- Loops & Control Flow ---

// FizzBuzz returns a slice of strings for numbers 1..n:
// multiples of 3 → "Fizz", multiples of 5 → "Buzz",
// multiples of both → "FizzBuzz", others → the number as string.
// TODO: loop 1..n, use fmt.Sprintf for number conversion
func FizzBuzz(n int) []string {
	return nil
}

// --- Errors ---

// ErrEmptyInput is returned when the input is empty.
var ErrEmptyInput = errors.New("empty input")

// FirstWord returns the first word of a sentence (split by spaces).
// Returns ErrEmptyInput if s is empty or contains only whitespace.
// TODO: use strings.Fields, check len
func FirstWord(s string) (string, error) {
	_ = strings.Fields // avoid unused import
	return "", nil
}
