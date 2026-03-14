# Exercise 06: Functions — Advanced

## Learning Objectives
- Use variadic functions (`...T`)
- Pass functions as arguments (higher-order functions)
- Return functions from functions (closures / function factories)

## Why This Matters (DevOps context)
Higher-order functions are used in Go middleware, retry logic, and functional options patterns. Variadic functions are common in logging libraries (e.g., `log.Printf`).

## Concept
Variadic parameters (`nums ...int`) accept zero or more arguments and behave like a slice inside the function. Functions are values — you can store, pass, and return them. A closure captures variables from its enclosing scope.

## Your Task
1. `Sum` — sum all variadic int arguments
2. `Apply` — call the provided function with x and return the result
3. `Multiplier` — return a closure that multiplies its argument by `factor`

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/06-functions-advanced/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
`nums` inside `Sum` is a `[]int`. Use a range loop to accumulate. For `Multiplier`, the returned function captures `factor`.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
func Multiplier(factor int) func(int) int {
    return func(n int) int {
        return n * factor
    }
}
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
`Sum`: `total := 0; for _, n := range nums { total += n }; return total`
</details>

## References
- [A Tour of Go — Variadic functions](https://go.dev/tour/moretypes/15) (see function values)
- [A Tour of Go — Function values](https://go.dev/tour/moretypes/24)
- [A Tour of Go — Closures](https://go.dev/tour/moretypes/25)
