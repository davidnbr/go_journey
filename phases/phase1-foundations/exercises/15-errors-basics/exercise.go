package exercise_15

import "errors"

// ErrDivisionByZero is a sentinel error for division by zero.
// TODO: define using errors.New
var ErrDivisionByZero = errors.New("")

// Divide returns a/b or ErrDivisionByZero if b == 0.
// TODO: check for zero divisor, return error
func Divide(a, b float64) (float64, error) {
	return 0, nil
}

// ParsePositive parses n and returns it only if > 0.
// Returns an error with message "must be positive: <n>" for non-positive n.
// TODO: use fmt.Errorf to format the error message
func ParsePositive(n int) (int, error) {
	return n, nil
}

// SafeSqrt returns the square root of n.
// Returns error "cannot take sqrt of negative number" for n < 0.
// For n >= 0, use math.Sqrt.
// TODO: guard negative input, return math.Sqrt(n)
func SafeSqrt(n float64) (float64, error) {
	return 0, nil
}
