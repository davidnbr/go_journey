# Exercise 09: Pointers — Basics

## Learning Objectives
- Take the address of a variable with `&`
- Dereference a pointer with `*`
- Distinguish between passing by value and by pointer
- Guard against nil pointer dereferences

## Why This Matters (DevOps context)
Pointer semantics are essential for modifying config structs in-place, implementing optional fields (nil = absent), and building efficient data structures. Every Go SDK uses `*string` for optional config fields.

## Concept
`&x` gives the address of `x`. `*p` gives the value at address `p`. `new(T)` allocates a zero-value `T` and returns `*T`. Dereferencing a nil pointer panics — always nil-check before dereferencing.

## Your Task
1. `Increment` — dereference and add 1
2. `Swap` — swap values via pointers
3. `NewInt` — allocate and return a pointer to a new int
4. `SafeDeref` — nil-safe dereference with a default

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/09-pointers-basics/
```

## Hints
<details><summary>Hint 1</summary>`*n++` or `*n = *n + 1` — both work.</details>
<details><summary>Hint 2</summary>`v := v; return &v` — takes address of local copy.</details>
<details><summary>Hint 3</summary>SafeDeref: `if ptr == nil { return defaultVal }; return *ptr`</details>

## References
- [A Tour of Go — Pointers](https://go.dev/tour/moretypes/1)
- [Effective Go — Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values)
