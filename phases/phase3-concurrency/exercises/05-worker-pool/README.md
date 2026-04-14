# Exercise 05: Worker Pool

## Learning Objectives
- Implement the worker pool pattern: fixed goroutines reading from a shared job channel
- Use `sync.WaitGroup` to close the results channel after all workers finish
- Understand why worker pools bound resource consumption

## Why This Matters (DevOps context)
Worker pools are used for parallel HTTP checks, database batch operations, and parallel resource provisioning. They prevent goroutine explosion when input is unbounded.

## Concept
One jobs channel, N worker goroutines all reading from it (Go's channel guarantees each job goes to exactly one worker). A WaitGroup tracks when all workers finish so we can close the results channel.

## Your Task
1. `Dispatch` — send all jobs to a channel in a goroutine, then close
2. `WorkerPool` — launch N workers; each reads jobs, applies fn, sends results; close results when all workers done

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/05-worker-pool/
```

## Hints
<details><summary>Hint 1</summary>

```go
var wg sync.WaitGroup
for i := 0; i < numWorkers; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        for job := range jobs {
            results <- Result{JobID: job.ID, Output: fn(job.Input)}
        }
    }()
}
go func() { wg.Wait(); close(results) }()
```
</details>

## References
- [Go Blog — Go Concurrency Patterns: Pipelines and cancellation](https://go.dev/blog/pipelines)
