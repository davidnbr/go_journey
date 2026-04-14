package exercise_02

import (
	"fmt"
	"strconv"
)

// ParseInt parses a string to int.
// Returns (0, error) for invalid strings.
// TODO: use strconv.Atoi
func ParseInt(s string) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

// SafeDivide divides a by b.
// Returns (0, error) if b == 0.
// TODO: guard zero divisor
func SafeDivide(a, b int) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

var _ = strconv.Atoi
