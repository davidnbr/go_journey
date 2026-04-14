package exercise_06

import "fmt"

// Stringify converts any value to a string.
// - If it's a string, return it directly.
// - If it's an int, return fmt.Sprintf("%d", v)
// - If it's a bool, return "true" or "false"
// - If it implements fmt.Stringer, call .String()
// - Otherwise return fmt.Sprintf("%v", v)
// TODO: use a type switch
func Stringify(v interface{}) string {
	return ""
}

// MustString asserts that v is a string and returns it.
// Panics with "expected string" if v is not a string.
// TODO: use a type assertion with ok check, then panic
func MustString(v interface{}) string {
	return ""
}

// AsInt attempts to extract an int from v.
// Returns (0, false) if v is not an int.
// TODO: use comma-ok type assertion
func AsInt(v interface{}) (int, bool) {
	return 0, false
}

var _ = fmt.Sprintf
