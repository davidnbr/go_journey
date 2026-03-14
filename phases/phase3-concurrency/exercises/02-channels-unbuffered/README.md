# Exercise 02: Channels — Unbuffered

## Learning Objectives
- Create and use unbuffered channels for synchronous communication
- Launch goroutines that produce values and close the channel
- Chain channels to form pipelines

## Why This Matters (DevOps context)
Unbuffered channels are synchronization points — sender blocks until receiver is ready. Used for signaling (done channels), streaming results, and goroutine coordination.

## Concept
`make(chan T)` creates an unbuffered channel. Send `ch <- v` blocks until a receiver is ready. `close(ch)` signals no more values; `range ch` loops until closed.

## Your Task
1. `Send` — goroutine sends 0..n-1 then closes
2. `Double` — goroutine transforms each value from `in`
3. `Collect` — drain channel into slice

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/02-channels-unbuffered/
```

## Hints
<details><summary>Hint 1</summary>Send: `ch := make(chan int); go func() { for i := 0; i < n; i++ { ch <- i }; close(ch) }(); return ch`</details>
<details><summary>Hint 2</summary>Double: `go func() { for v := range in { out <- v*2 }; close(out) }()`</details>
<details><summary>Hint 3</summary>Collect: `for v := range ch { result = append(result, v) }`</details>

## References
- [A Tour of Go — Channels](https://go.dev/tour/concurrency/2)
- [A Tour of Go — Range and close](https://go.dev/tour/concurrency/4)
