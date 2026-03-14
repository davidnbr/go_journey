package exercise_05

// Job represents a unit of work: an input int and a result int.
type Job struct {
	ID    int
	Input int
}

type Result struct {
	JobID  int
	Output int
}

// WorkerPool launches numWorkers goroutines that each read from jobs,
// apply fn to Job.Input, send Result to results, until jobs is closed.
// Closes results when all workers are done.
// TODO: launch numWorkers goroutines + WaitGroup to close results
func WorkerPool(numWorkers int, jobs <-chan Job, fn func(int) int) <-chan Result {
	results := make(chan Result)
	close(results) // stub
	return results
}

// Dispatch sends all jobs to the jobs channel and closes it.
// TODO: launch goroutine, send all jobs, close jobs
func Dispatch(jobs []Job) <-chan Job {
	ch := make(chan Job)
	close(ch) // stub: empty
	return ch
}
