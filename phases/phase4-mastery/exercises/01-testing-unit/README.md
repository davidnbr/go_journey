# Exercise 01: Testing — Unit Tests and Subtests

## Learning Objectives
- Write unit tests using the `testing` package
- Use `t.Run` for subtests (table-driven friendly)
- Understand test naming and failure messages

## Why This Matters (DevOps context)
Go's testing package is built-in — no test framework needed. Every Go open source project uses it. `go test ./...` runs all tests; `-v` shows subtest names; `-run TestFoo/subtest` runs specific subtests.

## Concept
`func TestXxx(t *testing.T)` is a test function. `t.Error` / `t.Errorf` mark failure without stopping. `t.Fatal` stops the test immediately. `t.Run("name", func(t *testing.T) {...})` creates a subtest.

## Your Task
1. Study `TestIsPalindrome` — `IsPalindrome` is already implemented; observe how subtests work
2. Implement `Fibonacci` — iterative approach
3. Implement `Clamp` — simple bounds clamping

## Run Tests
```bash
go test -v ./phases/phase4-mastery/exercises/01-testing-unit/
```

## Hints
<details><summary>Hint 1</summary>Fibonacci: `a, b := 0, 1; for i := 0; i < n; i++ { a, b = b, a+b }; return a`</details>
<details><summary>Hint 2</summary>Clamp: `if v < min { return min }; if v > max { return max }; return v`</details>

## References
- [pkg.go.dev — testing](https://pkg.go.dev/testing)
- [Go Blog — Using Subtests and Sub-benchmarks](https://go.dev/blog/subtests)
