# Exercise 07: Closures

## Learning Objectives
- Understand closure variable capture
- Build stateful functions without structs
- Implement the "once" initialization pattern

## Why This Matters (DevOps context)
Closures power Go's functional options pattern, connection pool initialization, and lazy resource loading. The `Once` pattern is `sync.Once` — used for singleton initialization in concurrent programs.

## Concept
A closure is a function that references variables from its enclosing scope. Those variables persist as long as the closure is reachable. Each call to `Counter()` creates an *independent* captured variable.

## Your Task
1. `Counter` — return a closure with its own `count` variable that increments each call
2. `Adder` — return a closure that adds `n` to its argument
3. `Once` — return a closure that calls `fn` exactly once, caching the result

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/07-closures/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
Declare a variable *before* the `return func()` statement. The returned function captures it by reference.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
func Counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
For `Once`: use `called bool` and `result string` captured vars. On first call set `called=true`, call `fn()`, store result. Return stored result on all calls.
</details>

## References
- [A Tour of Go — Closures](https://go.dev/tour/moretypes/25)
- [Go Blog — Functions as values](https://go.dev/blog/functions-as-values-and-closures)
