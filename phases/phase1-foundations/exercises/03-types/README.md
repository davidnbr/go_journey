# Exercise 03: Types

## Learning Objectives
- Perform explicit type conversions in Go
- Understand that Go has no implicit type coercion
- Use `len()` for string byte length

## Why This Matters (DevOps context)
Parsing config files, environment variables, and API responses constantly requires type conversion. Go's explicit approach prevents silent precision loss bugs common in dynamically typed languages.

## Concept
Go is statically typed with no implicit conversions. You must explicitly convert: `float64(myInt)`, `int(myFloat)`. `len(s)` returns byte count, not rune count — important for UTF-8 strings.

## Your Task
Implement the four conversion and inspection functions using explicit type conversion and `len()`.

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/03-types/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
Type conversion syntax: `TargetType(value)`. E.g., `float64(n)`.
</details>

<details><summary>Hint 2 — Pseudocode</summary>
`return float64(n)` / `return int(f)` / `return len(s)`
</details>

<details><summary>Hint 3 — Solution approach</summary>
`IsPositive`: `return n > 0` — comparison operators return bool directly.
</details>

## References
- [A Tour of Go — Type conversions](https://go.dev/tour/basics/13)
- [Go Spec — Conversions](https://go.dev/ref/spec#Conversions)
