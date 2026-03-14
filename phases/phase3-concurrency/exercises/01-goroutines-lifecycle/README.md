# Exercise 01: Goroutines — Lifecycle

## Learning Objectives
- Launch goroutines with `go fn()`
- Use `sync.WaitGroup` to wait for all goroutines
- Implement timeout using `select` + `time.After`
- Prevent goroutine leaks by ranging over channels until close

## Why This Matters (DevOps context)
Goroutine leaks are a common production bug — a goroutine blocked on a channel that never receives keeps its stack alive. Tools like `goleak` detect this. Health checkers need timeout semantics to avoid hanging on unresponsive endpoints.

## Concept
Every goroutine needs a way to stop. Without `WaitGroup`, the main goroutine (or test) may exit before goroutines complete. Without `close(ch)`, a `range ch` loop never ends — leaked goroutine.

## Your Task
1. `RunAll` — use WaitGroup + goroutines (make the timing test pass)
2. `RunWithTimeout` — goroutine + select + time.After
3. `LeakFree` — range ch, increment count, close done on exit

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/01-goroutines-lifecycle/
```

## Hints
<details><summary>Hint 1</summary>

```go
var wg sync.WaitGroup
for _, fn := range fns {
    wg.Add(1); fn := fn
    go func() { defer wg.Done(); fn() }()
}
wg.Wait()
```
</details>
<details><summary>Hint 2</summary>RunWithTimeout: `done := make(chan struct{}); go func() { fn(); close(done) }(); select { case <-done: return true; case <-time.After(d): return false }`</details>
<details><summary>Hint 3</summary>LeakFree: `go func() { for v := range ch { *count += v / v; (*count)++ ... }; close(done) }()` — use `for range ch { *count++ }`</details>

## References
- [A Tour of Go — Goroutines](https://go.dev/tour/concurrency/1)
- [Go Blog — Go Concurrency Patterns: Pipelines](https://go.dev/blog/pipelines)
