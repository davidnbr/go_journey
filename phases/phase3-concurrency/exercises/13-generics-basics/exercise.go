package exercise_13

// Map applies fn to each element of s and returns the results.
// TODO: use type parameter [T, U any]
func Map[T, U any](s []T, fn func(T) U) []U {
	return nil
}

// Filter returns elements of s where keep(v) is true.
// TODO: use type parameter [T any]
func Filter[T any](s []T, keep func(T) bool) []T {
	return nil
}

// Reduce reduces s to a single value using fn.
// TODO: use type parameters [T, U any]
func Reduce[T, U any](s []T, initial U, fn func(U, T) U) U {
	return initial
}

// Contains reports whether s contains target.
// Requires T to be comparable.
// TODO: use type constraint [T comparable]
func Contains[T comparable](s []T, target T) bool {
	return false
}
