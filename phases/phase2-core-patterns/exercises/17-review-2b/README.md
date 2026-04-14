# Exercise 17: Review 2b — Pointers, I/O, Defer, Modules

## Learning Objectives
Reinforce Phase 2 concepts 09–15 through three practical mini-implementations.

## Concepts Covered
- `io.Reader`/`io.Writer` pipeline
- Error wrapping with `%w`
- Functional options pattern (closure over pointer)

## Your Task
1. `Pipeline` — chain reader → transform → writer
2. `SafeDivide` — wrap ErrDivByZero
3. Functional options: `WithHost`, `WithPort`, `NewServerConfig`

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/17-review-2b/
```

## Hints
<details><summary>Hint 1</summary>Pipeline: read all, uppercase, write string to dst, return n=len(result), nil.</details>
<details><summary>Hint 2</summary>WithHost: `return func(c *ServerConfig) { c.Host = host }`</details>
<details><summary>Hint 3</summary>NewServerConfig: start with defaults, `for _, opt := range opts { opt(&cfg) }`, return cfg.</details>
