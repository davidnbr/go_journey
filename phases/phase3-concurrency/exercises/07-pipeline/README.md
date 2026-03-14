# Exercise 07: Pipeline Pattern

## Learning Objectives
- Build a composable channel pipeline (Generate → Filter → Map → Reduce)
- Each stage is a goroutine: input channel in, output channel out
- Understand how closing propagates through a pipeline

## Why This Matters (DevOps context)
Log processing pipelines, ETL jobs, and metric aggregation all follow this pattern. Go's channel pipelines are the foundation of tools like `kubectl logs` streaming and `docker events` processing.

## Concept
Each stage function: receives `<-chan T`, returns `<-chan T`, launches a goroutine internally. When the upstream closes, the `range` loop ends and the stage closes its own output — propagating through the pipeline.

## Your Task
1. `Generate` — produce integers start..end
2. `Filter` — pass through values where keep(v) is true
3. `Map` — transform each value
4. `Reduce` — synchronous reduction (no goroutine needed)

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/07-pipeline/
```

## Hints
<details><summary>Hint 1</summary>Each stage: `out := make(chan int); go func() { for v := range in { if/send }; close(out) }(); return out`</details>

## References
- [Go Blog — Go Concurrency Patterns: Pipelines](https://go.dev/blog/pipelines)
