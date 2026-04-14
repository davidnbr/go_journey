# Exercise 04: Constants

## Learning Objectives
- Declare typed and untyped constants
- Use `iota` to create enumerated values
- Implement a type-safe enumeration pattern

## Why This Matters (DevOps context)
Constants and iota are used for environment names (prod/staging/dev), log levels, status codes, and HTTP method sets. Go's typed constants prevent misuse at compile time.

## Concept
`const` values are compile-time immutable. `iota` auto-increments within a `const` block starting at 0. Untyped constants are flexible and can be used in expressions with any compatible type.

## Your Task
1. Fix `Pi` to equal `3.14159`
2. Fix `DaysInWeek` to equal `7`
3. The `Direction` iota enum is already correct — verify you understand why
4. Implement `DirectionName()` using a switch statement

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/04-constants/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
`iota` starts at 0 and increments by 1 for each const in the block. North=0, East=1, etc.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
switch d {
case North:
    return "North"
// ...
default:
    return "Unknown"
}
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
Change `const Pi = 0.0` to `const Pi = 3.14159` and `const DaysInWeek int = 0` to `const DaysInWeek int = 7`.
</details>

## References
- [A Tour of Go — Constants](https://go.dev/tour/basics/15)
- [Effective Go — Constants](https://go.dev/doc/effective_go#constants)
- [Go Spec — Iota](https://go.dev/ref/spec#Iota)
