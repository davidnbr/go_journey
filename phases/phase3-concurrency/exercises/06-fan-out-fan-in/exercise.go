package exercise_06

import "sync"

// FanOut distributes work from in to n output channels (round-robin).
// Each output channel gets roughly 1/n of the work.
// Closes all outputs when in is closed.
// TODO: create n channels, goroutine round-robins values
func FanOut(in <-chan int, n int) []<-chan int {
	return nil
}

// FanIn merges multiple input channels into one output channel.
// Closes output when all inputs are closed.
// TODO: one goroutine per input, WaitGroup to close output
func FanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	if len(inputs) == 0 {
		close(out)
		return out
	}
	var wg sync.WaitGroup
	// stub: close immediately
	close(out)
	_ = wg
	return out
}
