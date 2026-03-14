# Exercise 04: Select Statement

## Learning Objectives
- Use `select` to multiplex channel operations
- Implement non-blocking receives with `default`
- Implement timeouts with `time.After`
- Merge multiple channels into one

## Why This Matters (DevOps context)
`select` is the core of Go's reactive programming model. Health checkers use it for timeouts. Event processors use it to drain multiple input streams. Circuit breakers use it for fallback logic.

## Concept
`select` blocks until one of its cases can proceed, then executes that case. If multiple are ready, one is chosen at random. `default` makes it non-blocking. `time.After(d)` returns a channel that receives after d.

## Your Task
1. `First` — select on two channels, return first value
2. `WithDefault` — non-blocking receive
3. `Timeout` — select with time.After
4. `Merge` — combine two channels into one

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/04-select-statement/
```

## Hints
<details><summary>Hint 1</summary>`select { case v := <-ch1: return v; case v := <-ch2: return v }`</details>
<details><summary>Hint 2</summary>Merge: track open channels with nil; `if ch1 == nil && ch2 == nil { close(out); return }`. When a channel is drained, set it to nil.</details>

## References
- [A Tour of Go — Select](https://go.dev/tour/concurrency/5)
- [Go Blog — Go Concurrency Patterns: Timing out, moving on](https://go.dev/blog/concurrency-timeouts)
