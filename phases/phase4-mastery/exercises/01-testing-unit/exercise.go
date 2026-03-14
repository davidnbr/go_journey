package exercise_01

import "strings"

// IsPalindrome returns true if s is a palindrome (ignoring case and spaces).
func IsPalindrome(s string) bool {
	s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

// Fibonacci returns the nth Fibonacci number (0-indexed).
// Fib(0)=0, Fib(1)=1, Fib(2)=1, Fib(3)=2, ...
// TODO: implement iteratively
func Fibonacci(n int) int {
	return 0 // stub
}

// Clamp returns v clamped to [min, max].
// TODO: if v < min return min; if v > max return max; else v
func Clamp(v, min, max int) int {
	return 0 // stub
}
