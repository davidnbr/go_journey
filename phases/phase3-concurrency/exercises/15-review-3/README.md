# Exercise 15: Review 3 — Concurrency Patterns

## Learning Objectives
Combine goroutines, channels, context, and WaitGroup into the semaphore and bounded-parallelism patterns.

## Concepts Covered
- Semaphore: buffered channel as slot counter
- BoundedParallel: WaitGroup + Semaphore to limit concurrency
- Context propagation through concurrent operations

## Why This Matters (DevOps context)
Bounded parallelism prevents overwhelming downstream services. Rate-limited API callers, parallel Terraform plan operations with a concurrency limit, and kubectl bulk operations all use this exact pattern.

## Your Task
1. `NewSemaphore` — buffered channel with n slots
2. `Acquire`/`Release` — send/receive on the channel
3. `BoundedParallel` — WaitGroup + semaphore around each fn

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/15-review-3/
```

## Hints
<details><summary>Hint 1</summary>Semaphore: `ch: make(chan struct{}, n)`. Acquire: `s.ch <- struct{}{}`. Release: `<-s.ch`.</details>
<details><summary>Hint 2</summary>BoundedParallel: `sem.Acquire(); go func() { defer wg.Done(); defer sem.Release(); errs[i] = fn(ctx) }()`</details>

## References
- [Go Blog — Go Concurrency Patterns: Context](https://go.dev/blog/context)
