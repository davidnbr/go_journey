# Exercise 12: Race Conditions

## Learning Objectives
- Recognize data race patterns
- Fix races with `sync.Mutex`
- Use channels as an alternative to shared mutable state
- Run `go test -race` to detect races

## Why This Matters (DevOps context)
Data races are undefined behavior — they cause random crashes, silent corruption, and security vulnerabilities. The Go race detector finds them at runtime. CI should always include `-race`.

## Concept
A data race occurs when two goroutines access the same memory location concurrently and at least one is a write. `go test -race` instruments the binary to detect this.

## Your Task
1. Fix `SafeCounter.Inc` and `Value` with mutex locking
2. Implement `RaceFreeSumWithChannels` using a single goroutine (no mutex needed)

## Run Tests
```bash
go test -race ./phases/phase3-concurrency/exercises/12-race-conditions/
```

## Hints
<details><summary>Hint 1</summary>Same as exercise 09 — `c.mu.Lock(); defer c.mu.Unlock()`</details>
<details><summary>Hint 2</summary>RaceFreeSumWithChannels: `go func() { sum := 0; for v := range ch { sum += v }; result <- sum }()`</details>

## References
- [Go Blog — Introducing the Go Race Detector](https://go.dev/blog/race-detector)
