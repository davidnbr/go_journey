# Exercise 12: Package Design

## Learning Objectives
- Understand exported vs unexported identifiers (capitalization rule)
- Use unexported struct fields for encapsulation
- Write package-level constants and constructors

## Why This Matters (DevOps context)
Package encapsulation prevents callers from depending on internal state. Terraform providers, Kubernetes controllers, and CLI tools all use unexported fields to hide secrets, connections, and implementation details.

## Concept
In Go, identifiers starting with an uppercase letter are exported (public). Lowercase = unexported (package-private). This applies to types, functions, methods, and struct fields.

## Your Task
1. Add `secret string` as an unexported field to `Config`
2. `NewConfig` — return Config with host/port/secret populated
3. `IsValid` — check Host non-empty AND Port > 0
4. `GetVersion` — return the `Version` constant

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/12-packages-design/
```

## Hints
<details><summary>Hint 1</summary>Unexported fields: `secret string` (lowercase s) in the struct body.</details>
<details><summary>Hint 2</summary>Tests in the same package can access unexported fields — `c.secret`.</details>

## References
- [Effective Go — Names](https://go.dev/doc/effective_go#names)
- [Go Spec — Exported identifiers](https://go.dev/ref/spec#Exported_identifiers)
