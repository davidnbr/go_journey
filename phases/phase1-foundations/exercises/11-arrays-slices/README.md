# Exercise 11: Arrays and Slices

## Learning Objectives
- Understand the difference between fixed arrays and dynamic slices
- Use slice expressions `s[low:high]`
- Use `len()` and `cap()` to inspect slice state

## Why This Matters (DevOps context)
Slices are Go's primary sequence type — used for log lines, IP lists, configuration arrays, and batch job items. Understanding len vs cap prevents subtle `append` bugs.

## Concept
Arrays are value types with fixed size: `[5]int`. Slices are descriptors (pointer + len + cap) into an underlying array. `s[low:high]` creates a new slice header sharing the same backing array.

## Your Task
1. `FirstThree` — use a slice expression, handle short slices
2. `ArraySum` — sum a fixed `[5]int` array
3. `SliceInfo` — return len and cap
4. `RemoveLast` — drop the last element safely

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/11-arrays-slices/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
For FirstThree: `if len(s) < 3 { return s }; return s[:3]`.
For RemoveLast: guard empty, then `return s[:len(s)-1]`.
</details>

<details><summary>Hint 2 — Pseudocode</summary>
`length, capacity = len(s), cap(s)` — these are built-in functions.
</details>

<details><summary>Hint 3 — Solution approach</summary>
Arrays are passed by value — `ArraySum` receives a copy. Range over it normally.
</details>

## References
- [A Tour of Go — Arrays](https://go.dev/tour/moretypes/6)
- [A Tour of Go — Slices](https://go.dev/tour/moretypes/7)
- [Go Blog — Go Slices: usage and internals](https://go.dev/blog/slices-intro)
