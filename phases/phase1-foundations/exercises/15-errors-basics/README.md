# Exercise 15: Errors — Basics

## Learning Objectives
- Use the `error` interface as a return value
- Create sentinel errors with `errors.New`
- Format descriptive errors with `fmt.Errorf`
- Test for specific errors with `errors.Is`

## Why This Matters (DevOps context)
Go's explicit error returns are central to reliable infrastructure code. Every Terraform provider, every AWS SDK call, every file operation returns an error. Sentinel errors enable callers to handle specific failure modes without string matching.

## Concept
`error` is a built-in interface with one method: `Error() string`. Convention: return `nil` on success, a non-nil error on failure. `errors.New` creates simple errors. `fmt.Errorf` formats dynamic messages. `errors.Is` checks error identity.

## Your Task
1. Fix `ErrDivisionByZero` to have a meaningful message
2. `Divide` — return the sentinel error on zero divisor
3. `ParsePositive` — use `fmt.Errorf` for the error message
4. `SafeSqrt` — guard negative input

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/15-errors-basics/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
`var ErrDivisionByZero = errors.New("division by zero")`. In `Divide`: `if b == 0 { return 0, ErrDivisionByZero }`.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
import "fmt"
if n <= 0 {
    return 0, fmt.Errorf("must be positive: %d", n)
}
return n, nil
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
SafeSqrt: `import "math"`, then `return math.Sqrt(n), nil` for the happy path.
</details>

## References
- [A Tour of Go — Errors](https://go.dev/tour/methods/19)
- [Go Blog — Error handling and Go](https://go.dev/blog/error-handling-and-go)
- [pkg.go.dev — errors package](https://pkg.go.dev/errors)
