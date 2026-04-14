# Exercise 12: Slices — Advanced

## Learning Objectives
- Use `append` to build slices dynamically
- Use `copy` to make independent slice copies
- Combine maps and slices for deduplication

## Why This Matters (DevOps context)
Deduplication is essential for IP allow-lists, unique alert IDs, and batch deduplication pipelines. Deep vs shallow copy bugs are a common source of hard-to-find mutation bugs in Go services.

## Concept
`append` may or may not allocate a new backing array depending on capacity. To make a truly independent copy, use `make([]T, len(s))` + `copy(dst, src)`. A simple slice expression `s[:]` shares the backing array.

## Your Task
1. `Deduplicate` — remove duplicates preserving order (use a `map[int]bool`)
2. `Flatten` — concatenate a slice-of-slices (use `append(result, sub...)`)
3. `CopySlice` — return a deep copy that doesn't share the backing array

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/12-slices-advanced/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
Deduplicate: `seen := make(map[int]bool)`, skip if `seen[v]`, then set `seen[v] = true` and append.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
result := make([]int, len(s))
copy(result, s)
return result
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
Flatten: `append(result, sub...)` — the `...` unpacks the sub-slice into individual arguments.
</details>

## References
- [Go Blog — Slices internals](https://go.dev/blog/slices-intro)
- [A Tour of Go — Append](https://go.dev/tour/moretypes/15)
- [pkg.go.dev — builtin copy](https://pkg.go.dev/builtin#copy)
