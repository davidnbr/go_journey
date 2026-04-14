# Exercise 02: Variables

## Learning Objectives
- Understand Go's zero values for basic types
- Use `:=` (short declaration) vs `var`
- Use multiple return values

## Why This Matters (DevOps context)
Uninitialized variables in Go have predictable zero values — no null pointer surprises. This makes config parsing and flag defaults safe and explicit.

## Concept
Go variables have zero values: `0` for numeric types, `false` for bool, `""` for string, `nil` for pointers/slices/maps. Short declaration `:=` infers the type from the right-hand side.

## Your Task
1. `ZeroValues()` — return actual zero values (0, 0.0, false, "")
2. `Swap(a, b)` — return the two arguments in reversed order
3. `ShortDeclaration()` — use `:=` to assign x=10, y=20, return x+y

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/02-variables/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
Zero values are the defaults. For Swap, Go allows multiple return values: `return b, a`.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
func ZeroValues() (int, float64, bool, string) {
    return 0, 0.0, false, ""
}
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
For ShortDeclaration: `x := 10; y := 20; return x + y`
</details>

## References
- [A Tour of Go — Variables](https://go.dev/tour/basics/8)
- [A Tour of Go — Zero values](https://go.dev/tour/basics/12)
- [Effective Go — Variables](https://go.dev/doc/effective_go#variables)
