# Exercise 04: Interfaces — Basics

## Learning Objectives
- Define an interface with multiple methods
- Implement the interface implicitly (no `implements` keyword)
- Write functions that accept interface types for polymorphism

## Why This Matters (DevOps context)
Go's `io.Reader`, `io.Writer`, `http.Handler`, and `fmt.Stringer` are all interfaces. Building infrastructure tools against interfaces (not concrete types) makes them testable and extensible.

## Concept
Go interfaces are satisfied implicitly — any type with the required methods satisfies the interface. `Shape` requires `Area()` and `Perimeter()` — if your type has both, it IS a `Shape`.

## Your Task
1. Implement `Area()` and `Perimeter()` on `Circle`
2. Implement `Area()` and `Perimeter()` on `Rectangle`
3. Implement `TotalArea` to sum areas of any `[]Shape`

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/04-interfaces-basics/
```

## Hints
<details><summary>Hint 1</summary>`math.Pi * c.Radius * c.Radius` for circle area.</details>
<details><summary>Hint 2</summary>`2 * (r.Width + r.Height)` for rectangle perimeter.</details>
<details><summary>Hint 3</summary>`for _, s := range shapes { total += s.Area() }`</details>

## References
- [A Tour of Go — Interfaces](https://go.dev/tour/methods/9)
- [Effective Go — Interfaces](https://go.dev/doc/effective_go#interfaces)
