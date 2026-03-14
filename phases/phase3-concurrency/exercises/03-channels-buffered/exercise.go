package exercise_03

// Buffer sends all values into a buffered channel of size cap and returns it.
// Does NOT need a goroutine — the buffer allows sending without a receiver ready.
// TODO: make(chan int, cap), send all values, return (don't close)
func Buffer(values []int, cap int) chan int {
	return nil
}

// Batch reads from ch in batches of size n.
// Returns a slice of batches. Last batch may be smaller than n.
// TODO: read from ch, collect into batches
func Batch(ch <-chan int, n int) [][]int {
	return nil
}

// RateLimit returns a channel that receives one token per interval.
// Sends exactly count tokens then closes.
// Use time.Tick or time.NewTicker.
// TODO: goroutine: loop count times, block on ticker, send struct{}{}
func RateLimit(count int) <-chan struct{} {
	ch := make(chan struct{})
	close(ch)
	return ch
}
