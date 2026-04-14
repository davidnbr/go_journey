package exercise_07

// Counter returns a function that increments and returns a count starting at 0.
// Each call to the returned function should return 1, 2, 3, ...
// TODO: use a closure to capture a counter variable
func Counter() func() int {
	return nil
}

// Adder returns a function that adds n to its argument.
// TODO: return a closure capturing n
func Adder(n int) func(int) int {
	return nil
}

// Once returns a function that calls fn only on the first invocation.
// Subsequent calls are no-ops returning the first result.
// TODO: use a closure to track whether fn has been called
func Once(fn func() string) func() string {
	return nil
}
