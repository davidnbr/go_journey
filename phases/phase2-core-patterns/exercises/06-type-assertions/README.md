# Exercise 06: Type Assertions and Type Switches

## Learning Objectives
- Use type assertions `v.(T)` and the comma-ok form
- Use type switches to dispatch by concrete type
- Understand when to use `interface{}` (or `any`)

## Why This Matters (DevOps context)
JSON deserialization into `map[string]interface{}`, plugin systems, and dynamic configuration all require type assertions. Terraform's SDK uses type switches heavily for provider resource attributes.

## Concept
`v.(T)` panics if wrong type. `v, ok := x.(T)` is safe — sets ok=false without panic. Type switches handle multiple types cleanly: `switch v := x.(type) { case string: ... }`.

## Your Task
1. `Stringify` — type switch on common types + `fmt.Stringer` interface
2. `MustString` — assert with ok, panic if not string
3. `AsInt` — comma-ok type assertion

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/06-type-assertions/
```

## Hints
<details><summary>Hint 1</summary>In the type switch, `case fmt.Stringer:` matches any type implementing that interface.</details>
<details><summary>Hint 2</summary>`s, ok := v.(string); if !ok { panic("expected string") }; return s`</details>
<details><summary>Hint 3</summary>Order matters in type switch: put `string` before `fmt.Stringer` — `string` doesn't implement Stringer by default.</details>

## References
- [A Tour of Go — Type assertions](https://go.dev/tour/methods/15)
- [A Tour of Go — Type switches](https://go.dev/tour/methods/16)
