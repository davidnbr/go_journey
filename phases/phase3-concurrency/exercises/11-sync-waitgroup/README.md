# Exercise 11: sync.WaitGroup

## Learning Objectives
- Use `sync.WaitGroup` to run functions in parallel and collect ordered results
- Use position-indexed result slices (no mutex needed for separate positions)
- Collect errors per-function without losing order

## Why This Matters (DevOps context)
Parallel infrastructure checks — verifying N regions, N endpoints, N resources simultaneously — use this exact pattern. Results must come back in the same order as inputs for consistent reporting.

## Concept
Writing to separate slice positions from separate goroutines is safe — no mutex needed since different goroutines write different indices. `wg.Wait()` ensures all writes are complete before we read.

## Your Task
1. `Parallel` — run each fn in its own goroutine, store result at position i
2. `ParallelWithErrors` — same but collect errors

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/11-sync-waitgroup/
```

## Hints
<details><summary>Hint 1</summary>

```go
var wg sync.WaitGroup
for i, fn := range fns {
    wg.Add(1)
    i, fn := i, fn
    go func() { defer wg.Done(); results[i] = fn() }()
}
wg.Wait()
```
</details>

## References
- [pkg.go.dev — sync.WaitGroup](https://pkg.go.dev/sync#WaitGroup)
