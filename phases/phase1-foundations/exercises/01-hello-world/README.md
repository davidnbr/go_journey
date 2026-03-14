# Exercise 01: Hello, World!

## Learning Objectives
- Understand `package main` vs library packages
- Call a function and return a string value
- Run tests with `go test`

## Why This Matters (DevOps context)
Every CLI tool, health-check binary, or automation script starts here. Understanding how Go packages, compilation, and test execution work is the foundation for everything that follows.

## Concept
Go programs are organized into packages. `package main` is the entry point for executables. Library packages (like this exercise) export functions for use by other code. Functions return typed values explicitly.

## Your Task
Make `Greeting()` return `"Hello, World!"`.

## Run Tests
```bash
make check PHASE=phase1-foundations EX=01-hello-world
# or directly:
go test ./phases/phase1-foundations/exercises/01-hello-world/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
The function needs to return a string. Look at the return type in the signature.
</details>

<details><summary>Hint 2 — Pseudocode</summary>
`return "some string literal"`
</details>

<details><summary>Hint 3 — Solution approach</summary>
Change `return ""` to `return "Hello, World!"` — exact string including the comma and exclamation mark.
</details>

## References
- [A Tour of Go — Packages](https://go.dev/tour/basics/1)
- [Effective Go — Commentary](https://go.dev/doc/effective_go#commentary)
