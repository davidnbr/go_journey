# Exercise 08: unsafe and cgo — Awareness

## Learning Objectives
- Understand `unsafe.Sizeof`, `unsafe.Alignof` for memory layout inspection
- Know when `unsafe` is appropriate (performance-critical libraries, interop)
- Understand the basics of cgo (calling C from Go)

## Why This Matters (DevOps context)
Most DevOps Go code never needs `unsafe`. But understanding it helps you: debug memory layout issues, understand zero-copy string/byte conversions in high-throughput logs, and read library code that uses it.

## Concept
`unsafe.Sizeof(x)` returns the size in bytes of x's type — evaluated at compile time. `unsafe.Alignof(x)` returns the alignment. These are useful for understanding struct memory layout and padding.

## Your Task
1. `SizeOf` — return `int(unsafe.Sizeof(v))` (note: for interface{}, this is the size of the interface, not the contained value)
2. `AlignOf` — return `int(unsafe.Alignof(v))`
3. `StringToBytes` — implement with the safe `[]byte(s)` copy (the stub is already correct)

## Run Tests
```bash
go test ./phases/phase4-mastery/exercises/08-unsafe-cgo/
```

## Note
The test for SizeOf passes an `interface{}` — `unsafe.Sizeof` on interface{} returns the size of the interface header (16 bytes on 64-bit), not the contained value. This is a known nuance. For the bool test, pass the bool directly via reflect or use a typed variable.

## Hints
<details><summary>Hint 1</summary>`return int(unsafe.Sizeof(v))` — but the bool test may not work as expected with interface{} wrapper. That's the lesson.</details>

## References
- [pkg.go.dev — unsafe](https://pkg.go.dev/unsafe)
- [Go Blog — cgo](https://go.dev/blog/cgo)
