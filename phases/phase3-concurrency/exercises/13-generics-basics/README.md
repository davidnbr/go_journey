# Exercise 13: Generics — Basics (Go 1.18+)

## Learning Objectives
- Write generic functions using type parameters `[T any]`
- Use the `comparable` constraint for equality checks
- Implement the classic `Map`, `Filter`, `Reduce`, `Contains` as type-safe generics

## Why This Matters (DevOps context)
Generics eliminate repetitive code for common data transformations. Instead of `MapStrings`, `MapInts`, `MapErrors` — one `Map[T, U]` handles all. Used in Terraform provider SDKs and Kubernetes controller helpers.

## Concept
`func Foo[T any](s []T) []T` — T is a type parameter inferred at call site. `[T comparable]` restricts T to types that support `==`. `[T, U any]` uses two type parameters for different input/output types.

## Your Task
Implement `Map`, `Filter`, `Reduce`, and `Contains` using generics. No casting, no `interface{}`.

## Run Tests
```bash
go test ./phases/phase3-concurrency/exercises/13-generics-basics/
```

## Hints
<details><summary>Hint 1</summary>Map: `result := make([]U, len(s)); for i, v := range s { result[i] = fn(v) }; return result`</details>
<details><summary>Hint 2</summary>Contains: `for _, v := range s { if v == target { return true } }; return false`</details>

## References
- [A Tour of Go — Generics](https://go.dev/tour/generics/1)
- [Go Blog — An Introduction To Generics](https://go.dev/blog/intro-generics)
