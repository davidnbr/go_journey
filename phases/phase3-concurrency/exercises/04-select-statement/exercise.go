package exercise_04

import "time"

// First returns the first value received from either ch1 or ch2.
// TODO: use select to receive from whichever is ready first
func First(ch1, ch2 <-chan string) string {
	return ""
}

// WithDefault receives from ch if a value is available, otherwise returns defaultVal.
// Must not block.
// TODO: use select with default case
func WithDefault(ch <-chan int, defaultVal int) int {
	return 0
}

// Timeout reads from ch but returns ("", true) if nothing arrives within d.
// Returns (value, false) on success.
// TODO: select with ch and time.After(d)
func Timeout(ch <-chan string, d time.Duration) (string, bool) {
	return "", false
}

// Merge merges two input channels into one output channel.
// Closes output when both inputs are closed.
// TODO: goroutine with select, track which channels are open
func Merge(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	close(out)
	return out
}
