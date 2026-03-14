package exercise_14

import "strings"

// IsPalindrome returns true if s is a palindrome (case-insensitive, ASCII only).
// TODO: compare string to its reverse; use strings.ToLower
func IsPalindrome(s string) bool {
	return false
}

// CountRunes returns the number of Unicode code points (runes) in s.
// This differs from len(s) for multi-byte characters.
// TODO: use range over string to count runes
func CountRunes(s string) int {
	return 0
}

// ReverseString returns s with its rune order reversed.
// Must handle multi-byte Unicode correctly.
// TODO: convert to []rune, reverse, convert back
func ReverseString(s string) string {
	return ""
}

// ContainsAll returns true if s contains all substrings in subs.
// TODO: use strings.Contains in a loop
func ContainsAll(s string, subs []string) bool {
	return false
}

// avoid unused import error — strings is used by the student's implementation
var _ = strings.Contains
