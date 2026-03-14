# Exercise 08: Control Flow — if

## Learning Objectives
- Use `if`, `else if`, `else` chains
- Understand the if-with-initializer syntax
- Implement range validation with early returns

## Why This Matters (DevOps context)
Conditional logic is the backbone of configuration validation, threshold alerting (CPU > 90% → alert), and health check categorization.

## Concept
Go's `if` requires no parentheses around the condition. The if-init form (`if x := expr; condition`) scopes `x` to the if/else block, keeping scope tight.

## Your Task
1. `Classify` — three-way classification using if/else if/else
2. `AbsoluteValue` — simple if/else, or use a comparison
3. `ParseScore` — range validation + grade assignment

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/08-control-if/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
Check bounds first (`< 0` or `> 100`) before grade comparisons. Return "invalid" early.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
if score < 0 || score > 100 { return "invalid" }
if score >= 90 { return "A" }
...
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
`AbsoluteValue`: `if n < 0 { return -n }; return n`
</details>

## References
- [A Tour of Go — If](https://go.dev/tour/flowcontrol/5)
- [A Tour of Go — If with short statement](https://go.dev/tour/flowcontrol/6)
- [Effective Go — If](https://go.dev/doc/effective_go#if)
