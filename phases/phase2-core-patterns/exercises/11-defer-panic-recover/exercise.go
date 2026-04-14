package exercise_11

import "fmt"

// WithCleanup runs fn and always calls cleanup, even if fn panics.
// Returns the error from fn (or nil). Does NOT suppress panics.
// TODO: use defer to call cleanup()
func WithCleanup(fn func() error, cleanup func()) (err error) {
	return fmt.Errorf("not implemented")
}

// SafeRun runs fn and recovers from any panic.
// Returns the panic value as an error string, or nil if no panic.
// TODO: defer a recover() that captures the panic value
func SafeRun(fn func()) (err error) {
	return nil
}

// DeferOrder appends to log in defer order (LIFO).
// Appends "first", "second", "third" via defer, returns the log.
// TODO: defer three appends; they execute in reverse (LIFO) order
func DeferOrder() []string {
	var log []string
	// TODO: defer three append calls
	return log
}
