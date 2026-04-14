package exercise_02

// Send sends values 0..n-1 on the returned channel, then closes it.
// TODO: create chan int, launch goroutine to send+close
func Send(n int) <-chan int {
	ch := make(chan int)
	close(ch) // stub: immediately closed empty channel
	return ch
}

// Double reads from in, doubles each value, sends to returned channel.
// Closes output when in is closed.
// TODO: create out chan, goroutine: range in, send v*2, close out
func Double(in <-chan int) <-chan int {
	out := make(chan int)
	close(out) // stub
	return out
}

// Collect drains ch into a slice and returns it.
// TODO: range over ch, append
func Collect(ch <-chan int) []int {
	return nil
}
