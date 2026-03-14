# Exercise 15: Goroutines — Preview

## Learning Objectives
- Launch goroutines with the `go` keyword
- Use `sync.WaitGroup` to wait for all goroutines to complete
- Understand that goroutines are cheap, concurrent, not parallel by default

## Why This Matters (DevOps context)
Concurrent health checks, parallel resource provisioning, and fan-out patterns are the core of Go's concurrency model. This exercise introduces the primitive — Phase 3 dives deep.

## Concept
`go fn()` launches `fn` as a goroutine — it runs concurrently with the caller. `sync.WaitGroup` coordinates: `wg.Add(1)` before the goroutine, `defer wg.Done()` inside, `wg.Wait()` to block until all done.

## Your Task
Modify `RunConcurrently` to actually run the functions concurrently using goroutines and a WaitGroup. The current stub runs them sequentially — make it concurrent.

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/15-goroutines-preview/
```

## Hints
<details><summary>Hint 1</summary>

```go
var wg sync.WaitGroup
for _, fn := range fns {
    wg.Add(1)
    fn := fn  // capture loop variable
    go func() {
        defer wg.Done()
        fn()
    }()
}
wg.Wait()
```
</details>

## References
- [A Tour of Go — Goroutines](https://go.dev/tour/concurrency/1)
- [A Tour of Go — WaitGroups](https://pkg.go.dev/sync#WaitGroup)
