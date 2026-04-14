# Exercise 06: Fan-Out / Fan-In

## Learning Objectives
- Distribute work across multiple goroutines (fan-out)
- Merge results from multiple goroutines into one channel (fan-in)
- Use `sync.WaitGroup` to close the merged channel cleanly

## Why This Matters (DevOps context)
Fan-out/fan-in is the basis of parallel data processing pipelines: fan-out to N workers, fan-in to collect results. Used in parallel test runners, log aggregators, and multi-region health checks.

## Concept
Fan-in: launch one goroutine per input channel, each forwarding to the shared output. WaitGroup tracks when all are done. Fan-out: distribute values round-robin (or randomly) to N outputs.

## Your Task
1. `FanIn` — merge N channels, close output when all done
2. `FanOut` — distribute input to N channels round-robin

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/06-fan-out-fan-in/
```

## Hints
<details><summary>Hint 1</summary>FanIn: `for _, ch := range inputs { wg.Add(1); go func(c <-chan int) { defer wg.Done(); for v := range c { out <- v } }(ch) }`</details>
<details><summary>Hint 2</summary>FanOut: `i % n` selects which output channel to send to.</details>

## References
- [Go Blog — Pipelines and cancellation](https://go.dev/blog/pipelines)
