# Exercise 13: Maps

## Learning Objectives
- Create and populate maps with `make`
- Use the comma-ok idiom for safe key lookups
- Iterate over maps with `range`

## Why This Matters (DevOps context)
Maps are the go-to structure for config key-value pairs, frequency counters, tag sets, and service registries. The comma-ok pattern prevents silent zero-value returns.

## Concept
Map declaration: `make(map[K]V)`. Access: `m[key]` returns zero value if missing — use `v, ok := m[key]` to distinguish "missing" from "stored zero". Maps are reference types; `nil` maps panic on write.

## Your Task
1. `WordCount` — count word frequencies; return empty map (not nil) for empty input
2. `Lookup` — wrap the comma-ok idiom
3. `Invert` — swap keys and values

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/13-maps/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
Initialize: `result := make(map[string]int)`. Then `result[word]++` — this works even on first access (zero value is 0).
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
v, ok := m[key]
return v, ok
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
Invert: `for k, v := range m { result[v] = k }`.
</details>

## References
- [A Tour of Go — Maps](https://go.dev/tour/moretypes/19)
- [Effective Go — Maps](https://go.dev/doc/effective_go#maps)
- [Go Blog — Maps in action](https://go.dev/blog/maps)
