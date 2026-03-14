# Exercise 02: Structs — Value vs Pointer Receivers

## Learning Objectives
- Choose between value and pointer receivers correctly
- Understand that pointer receivers are needed to mutate state
- Implement the `fmt.Stringer` interface via `String() string`

## Why This Matters (DevOps context)
Most Go services use pointer receivers for anything with state (connection pools, config structs, HTTP handlers). Mixing receiver types on the same type causes confusing bugs and vet warnings.

## Concept
Value receiver (`func (c Circle)`) — receives a copy; mutations don't affect the original. Pointer receiver (`func (c *Circle)`) — receives the address; mutations affect the original. Rule of thumb: if any method mutates state, make all methods pointer receivers.

## Your Task
1. `Area` — value receiver, return area formula
2. `Scale` — pointer receiver, multiply radius in-place
3. `String` — value receiver, return `"Circle(r=<radius>)"`

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/02-structs-methods/
```

## Hints
<details><summary>Hint 1</summary>`func (c *Circle) Scale(factor float64) { c.Radius *= factor }`</details>
<details><summary>Hint 2</summary>`fmt.Sprintf("Circle(r=%g)", c.Radius)` — `%g` removes trailing zeros.</details>
<details><summary>Hint 3</summary>Test expects exactly `"Circle(r=3)"` for radius 3.</details>

## References
- [A Tour of Go — Pointer receivers](https://go.dev/tour/methods/4)
- [Effective Go — Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values)
