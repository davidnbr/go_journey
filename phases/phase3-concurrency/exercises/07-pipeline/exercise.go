package exercise_07

// Generate produces integers from start to end (inclusive) on a channel.
// TODO: goroutine, send start..end, close
func Generate(start, end int) <-chan int {
	ch := make(chan int)
	close(ch)
	return ch
}

// Filter removes values from in where keep(v) is false.
// TODO: goroutine, range in, if keep(v) send to out, close out
func Filter(in <-chan int, keep func(int) bool) <-chan int {
	out := make(chan int)
	close(out)
	return out
}

// Map transforms each value from in using transform.
// TODO: goroutine, range in, send transform(v), close out
func Map(in <-chan int, transform func(int) int) <-chan int {
	out := make(chan int)
	close(out)
	return out
}

// Reduce collects all values from in and reduces them with fn.
// initial is the starting accumulator value.
// TODO: range in, apply fn(acc, v), return final acc
func Reduce(in <-chan int, initial int, fn func(acc, v int) int) int {
	return initial
}
