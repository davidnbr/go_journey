package exercise_03

import "strings"

// ConcatLoop concatenates n copies of s using a loop.
// TODO: use a simple loop with += (intentionally slow)
func ConcatLoop(s string, n int) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}
	return result
}

// ConcatBuilder concatenates n copies of s using strings.Builder.
// TODO: use strings.Builder for O(n) performance
func ConcatBuilder(s string, n int) string {
	return strings.Repeat(s, 0) // stub: returns ""
}

// ConcatRepeat uses strings.Repeat.
// TODO: return strings.Repeat(s, n)
func ConcatRepeat(s string, n int) string {
	return "" // stub
}
