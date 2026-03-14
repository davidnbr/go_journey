# Exercise 10: Loops

## Learning Objectives
- Use Go's single loop construct `for` in all its forms
- Iterate over slices with `range`
- Build slices incrementally with `append`

## Why This Matters (DevOps context)
Polling loops, retry loops, batch processors, and sequence generators are everyday infrastructure patterns. Go's unified `for` loop handles all of them.

## Concept
Go has one loop keyword: `for`. Forms: `for i := 0; i < n; i++` (C-style), `for condition` (while-style), `for range` (iteration). `break` and `continue` work as expected.

## Your Task
1. `SumUpTo` — sum 1..n using a for loop
2. `CollectEven` — filter even numbers using range + append
3. `Fibonacci` — build the sequence iteratively

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/10-loops/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
For CollectEven: `if num%2 == 0 { result = append(result, num) }`.
For Fibonacci: track `a, b := 0, 1` and shift on each iteration.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
a, b := 0, 1
for i := 0; i < n; i++ {
    result = append(result, a)
    a, b = b, a+b
}
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
SumUpTo: `for i := 1; i <= n; i++ { sum += i }` — remember to guard `n < 1`.
</details>

## References
- [A Tour of Go — For](https://go.dev/tour/flowcontrol/1)
- [A Tour of Go — Range](https://go.dev/tour/moretypes/22)
- [Effective Go — For](https://go.dev/doc/effective_go#for)
