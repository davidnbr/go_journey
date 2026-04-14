# Exercise 09: Control Flow — switch

## Learning Objectives
- Use `switch` for multi-branch dispatch (cleaner than long if/else chains)
- Group multiple cases with a comma list
- Use type switch to inspect `interface{}` values

## Why This Matters (DevOps context)
HTTP status routing, log level dispatch, and cloud event type handling are all switch-shaped problems. Type switches power generic serializers and plugin dispatchers.

## Concept
Go's `switch` does not fall through by default (use `fallthrough` explicitly). Multiple values in one case: `case "Sat", "Sun":`. Type switch: `switch v := x.(type) { case int: ... }`.

## Your Task
1. `DayType` — map day names to Weekday/Weekend/Unknown
2. `HTTPStatus` — map status codes to descriptions
3. `Describe` — use a type switch on `interface{}`

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/09-control-switch/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
Group Monday–Friday in one case using commas. Use `default:` for Unknown/unrecognized values.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    return "Weekday"
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
Type switch: `switch v := val.(type) { case int: return "int"; case string: return "string"; ... }`
Note: `v` is available with the concrete type inside each case.
</details>

## References
- [A Tour of Go — Switch](https://go.dev/tour/flowcontrol/9)
- [A Tour of Go — Type switches](https://go.dev/tour/methods/16)
- [Effective Go — Switch](https://go.dev/doc/effective_go#switch)
