# Exercise 03: Channels — Buffered

## Learning Objectives
- Create buffered channels with `make(chan T, n)`
- Understand that buffered sends don't block until the buffer is full
- Implement batching from a channel stream

## Why This Matters (DevOps context)
Buffered channels decouple producers from consumers — essential for rate limiting, batch processing, and preventing fast producers from blocking on slow consumers (backpressure).

## Concept
`make(chan T, n)` creates a buffered channel with capacity n. Sends succeed without a receiver until the buffer fills. `len(ch)` = current items, `cap(ch)` = capacity.

## Your Task
1. `Buffer` — fill a buffered channel synchronously (no goroutine needed)
2. `Batch` — collect channel values into fixed-size batches
3. `RateLimit` — emit N tokens with pacing

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/03-channels-buffered/
```

## Hints
<details><summary>Hint 1</summary>Buffer: `ch := make(chan int, cap); for _, v := range values { ch <- v }; return ch`</details>
<details><summary>Hint 2</summary>Batch: accumulate into `current []int`; when `len(current) == n`, append to batches and reset.</details>

## References
- [A Tour of Go — Buffered Channels](https://go.dev/tour/concurrency/3)
