package taskqueue

import (
	"context"
	"fmt"
	"sync"
)

// Task is a unit of work.
type Task struct {
	ID int
	Fn func(ctx context.Context) error
}

// TaskResult holds the outcome of a task.
type TaskResult struct {
	TaskID int
	Error  error
}

// Queue is a concurrent task queue with bounded workers.
type Queue struct {
	workers int
	jobs    chan Task
	results chan TaskResult
	wg      sync.WaitGroup
	once    sync.Once
}

// NewQueue creates a Queue with the given number of workers and buffer size.
// TODO: initialize jobs channel, results channel, start workers
func NewQueue(workers, bufferSize int) *Queue {
	return &Queue{}
}

// Start launches the worker goroutines.
// TODO: start workers goroutines, each reading from q.jobs
func (q *Queue) Start(ctx context.Context) {
	// no-op stub
}

// Submit adds a task to the queue.
// Returns error if the queue is full or stopped.
// TODO: send task to q.jobs (non-blocking with default)
func (q *Queue) Submit(t Task) error {
	return fmt.Errorf("not implemented")
}

// Results returns the results channel for reading outcomes.
func (q *Queue) Results() <-chan TaskResult {
	return q.results
}

// Stop signals no more tasks and waits for all workers to finish.
// TODO: close q.jobs, wg.Wait(), close q.results (sync.Once)
func (q *Queue) Stop() {
	// no-op stub
}
