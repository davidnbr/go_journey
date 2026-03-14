# Exercise 05: Functions — Basics

## Learning Objectives
- Define functions with parameters and return values
- Return multiple values (quotient + success flag)
- Use named return values

## Why This Matters (DevOps context)
Go functions returning `(value, error)` or `(value, bool)` are the foundation of all Go error handling. Every AWS SDK call, every file read, every HTTP request follows this pattern.

## Concept
Functions are first-class in Go. Multiple return values replace out-parameters and exceptions. Named return values (declared in the signature) are initialized to zero values.

## Your Task
1. `Add` — return sum of two ints
2. `Divide` — return quotient and success bool; handle zero divisor
3. `MinMax` — return minimum and maximum of a slice; handle empty slice

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/05-functions-basics/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
For Divide: check `if b == 0` first. For MinMax: initialize min/max to the first element, then iterate.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
if len(nums) == 0 { return 0, 0 }
min, max = nums[0], nums[0]
for _, n := range nums { ... }
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
Use `if n < min { min = n }` and `if n > max { max = n }` inside the range loop.
</details>

## References
- [A Tour of Go — Functions](https://go.dev/tour/basics/4)
- [A Tour of Go — Multiple results](https://go.dev/tour/basics/6)
- [A Tour of Go — Named return values](https://go.dev/tour/basics/7)
