# Exercise 09: Performance Profiling

## Learning Objectives
- Compare O(n^2) vs O(1) and O(n) vs O(log n) algorithms via benchmarks
- Use `go test -bench=. -benchmem -cpuprofile=cpu.prof` for profiling
- Understand escape analysis with `go build -gcflags="-m"`

## Why This Matters (DevOps context)
Performance matters in hot paths: log parsing, metric aggregation, health check loops. Benchmarks catch regressions; pprof identifies bottlenecks. The difference between O(n) and O(log n) is observable at 10k+ elements.

## Your Task
1. `SumSquares` — naive loop (O(n))
2. `SumSquaresFast` — closed-form formula (O(1))
3. `ContainsLinear` — linear scan (O(n))
4. `ContainsBinary` — binary search (O(log n)) using `sort.SearchInts`

## Run Tests and Benchmarks
```bash
go test ./phases/phase4-mastery/exercises/09-performance-profiling/
go test -bench=. -benchmem ./phases/phase4-mastery/exercises/09-performance-profiling/
```

## Hints
<details><summary>Hint 1</summary>SumSquaresFast: `return n * (n + 1) * (2*n + 1) / 6`</details>
<details><summary>Hint 2</summary>ContainsBinary: `i := sort.SearchInts(s, target); return i < len(s) && s[i] == target`</details>

## References
- [pkg.go.dev — sort.SearchInts](https://pkg.go.dev/sort#SearchInts)
- [Go Blog — Profiling Go Programs](https://go.dev/blog/profiling-go-programs)
