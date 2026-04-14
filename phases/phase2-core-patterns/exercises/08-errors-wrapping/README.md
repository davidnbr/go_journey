# Exercise 08: Errors — Wrapping

## Learning Objectives
- Wrap errors with `fmt.Errorf("context: %w", err)` (Go 1.13+)
- Use `errors.Is` to check identity through a chain
- Use `errors.As` to extract typed errors through a chain

## Why This Matters (DevOps context)
Error wrapping is how Go stack traces work without a call stack — each layer adds context. AWS SDK, Terraform, and kubectl all wrap errors to preserve the root cause while adding operation context.

## Concept
`%w` in `fmt.Errorf` wraps the error; `errors.Is(err, target)` unwraps the chain. `errors.Unwrap` returns the inner error. Wrapping preserves the original error identity for programmatic checks.

## Your Task
1. `OpenConfig` — return wrapped `ErrNotFound` or `ErrPermission` with context strings
2. `LoadConfig` — call `OpenConfig` and add a `"loadConfig: "` wrapper

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/08-errors-wrapping/
```

## Hints
<details><summary>Hint 1</summary>`fmt.Errorf("opening %s: %w", name, ErrNotFound)`</details>
<details><summary>Hint 2</summary>`errors.Is` unwraps automatically through the whole chain.</details>
<details><summary>Hint 3</summary>LoadConfig: `if err := OpenConfig(name); err != nil { return fmt.Errorf("loadConfig: %w", err) }`</details>

## References
- [Go Blog — Working with errors in Go 1.13](https://go.dev/blog/go1.13-errors)
- [pkg.go.dev — fmt.Errorf](https://pkg.go.dev/fmt#Errorf)
