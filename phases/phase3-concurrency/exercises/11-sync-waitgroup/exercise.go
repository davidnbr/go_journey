package exercise_11

import (
	"sync"
)

// Parallel runs all fns in parallel and collects results.
// Returns results in the SAME ORDER as fns (not completion order).
// TODO: WaitGroup + slice indexed by position
func Parallel(fns []func() int) []int {
	results := make([]int, len(fns))
	for i, fn := range fns {
		results[i] = fn() // stub: sequential
	}
	return results
}

// ParallelWithErrors runs all fns in parallel.
// Returns all errors (nil for successes); preserves order.
// TODO: WaitGroup + indexed error slice
func ParallelWithErrors(fns []func() error) []error {
	errs := make([]error, len(fns))
	for i, fn := range fns {
		errs[i] = fn() // stub: sequential
	}
	return errs
}

var _ = sync.WaitGroup{}
