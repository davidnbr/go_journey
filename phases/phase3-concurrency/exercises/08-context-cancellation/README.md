# Exercise 08: Context — Cancellation and Timeouts

## Learning Objectives
- Use `context.WithCancel` and `context.WithTimeout` to control goroutine lifecycles
- Check `ctx.Done()` in long-running operations
- Always `defer cancel()` to prevent context leaks

## Why This Matters (DevOps context)
`context.Context` is the standard mechanism for cancellation in Go services. HTTP handlers, database queries, gRPC calls, and AWS SDK operations all accept a context. Without cancellation, a slow downstream causes goroutine accumulation.

## Concept
`ctx.Done()` returns a channel closed when the context is cancelled/timed out. `ctx.Err()` returns the reason. Pattern: `select { case <-time.After(d): return nil; case <-ctx.Done(): return ctx.Err() }`.

## Your Task
1. `SleepWithContext` — respect cancellation
2. `FetchAll` — skip remaining fetches when context is done
3. `WithDeadline` — wrap a function with a timeout context

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/08-context-cancellation/
```

## Hints
<details><summary>Hint 1</summary>`select { case <-time.After(d): return nil; case <-ctx.Done(): return ctx.Err() }`</details>
<details><summary>Hint 2</summary>WithDeadline: `ctx, cancel := context.WithTimeout(context.Background(), timeout); defer cancel(); return fn(ctx)`</details>

## References
- [pkg.go.dev — context package](https://pkg.go.dev/context)
- [Go Blog — Contexts and structs](https://go.dev/blog/context-and-structs)
