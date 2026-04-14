# Exercise 16: Review 1 — Spaced Repetition

## Learning Objectives
Reinforce Phase 1 concepts without referencing earlier exercises. If you need to look anything up, that's a signal to review that exercise before continuing to Phase 2.

## Concepts Covered
- Type conversion and arithmetic (`CelsiusToFahrenheit`)
- Closures with map state (`RateCounter`)
- Loops + slices + maps together (`GroupByLength`)
- Control flow + string formatting (`FizzBuzz`)
- Error returns + sentinel errors (`FirstWord`)

## Why This Matters (DevOps context)
A monitoring agent needs all of these simultaneously: temperature conversion for sensor data, rate counters for metrics, grouping events by type, FizzBuzz-style modulo logic for rotation schedules, and error-safe string parsing.

## Your Task
Implement all five functions without looking back at previous exercises. Use hints only if you're truly stuck.

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/16-review-1/
```

## Hints
<details><summary>Hint 1 — CelsiusToFahrenheit</summary>
`return c*9/5 + 32` — be careful: use float64 literals or ensure no integer division.
</details>

<details><summary>Hint 2 — RateCounter</summary>
`counts := make(map[string]int)` captured in the closure. For `""`, sum all values.
</details>

<details><summary>Hint 3 — FizzBuzz</summary>
Check `n%15 == 0` first (FizzBuzz), then `n%3`, then `n%5`, else `fmt.Sprintf("%d", n)`.
</details>

## References
- Review: exercises 01–15 as needed
