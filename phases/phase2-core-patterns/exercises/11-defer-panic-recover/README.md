# Exercise 11: Defer, Panic, and Recover

## Learning Objectives
- Use `defer` to guarantee cleanup (file close, mutex unlock, connection release)
- Understand LIFO defer ordering
- Use `recover()` to convert panics into errors

## Why This Matters (DevOps context)
`defer` is used for every resource release in Go: `defer file.Close()`, `defer mu.Unlock()`, `defer rows.Close()`. Panic/recover is how HTTP servers stay alive after handler panics — the standard library's `http.Server` uses this.

## Concept
`defer f()` schedules `f` to run when the enclosing function returns, even on panic. Multiple defers run in LIFO order. `recover()` inside a deferred function catches a panic and returns the panic value.

## Your Task
1. `WithCleanup` — defer cleanup(), then call fn() and return its error
2. `SafeRun` — defer a recover() that sets err if there's a panic
3. `DeferOrder` — demonstrate LIFO by deferring three appends

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/11-defer-panic-recover/
```

## Hints
<details><summary>Hint 1</summary>Named return `err` lets the deferred function access it: `defer func() { cleanup() }()`</details>
<details><summary>Hint 2</summary>SafeRun: `defer func() { if r := recover(); r != nil { err = fmt.Errorf("%v", r) } }()`</details>
<details><summary>Hint 3</summary>DeferOrder: defer log = append(log, "first"), then "second", then "third" — they run in reverse.</details>

## References
- [A Tour of Go — Defer](https://go.dev/tour/flowcontrol/12)
- [Go Blog — Defer, Panic, and Recover](https://go.dev/blog/defer-panic-and-recover)
