# Exercise 03: Testing — Benchmarks

## Learning Objectives
- Write `Benchmark*` functions using `*testing.B`
- Interpret benchmark output (ns/op, allocs/op)
- Compare performance of different implementations

## Why This Matters (DevOps context)
Benchmarks catch performance regressions in CI. A 10x slowdown in a hot path (serialization, validation) is detectable before deployment. `pprof` profiles show where time is spent.

## Concept
`func BenchmarkFoo(b *testing.B) { for i := 0; i < b.N; i++ { ... } }`. Run with `go test -bench=. -benchmem`. `b.N` is auto-tuned. `strings.Builder` avoids quadratic allocations vs `+=`.

## Your Task
1. `ConcatLoop` is already implemented (slow: O(n²) allocations) — run the benchmark to see
2. Implement `ConcatBuilder` using `strings.Builder`
3. Implement `ConcatRepeat` using `strings.Repeat`

## Run Tests and Benchmarks
```bash
go test ./phases/phase4-mastery/exercises/03-testing-benchmarks/
go test -bench=. -benchmem ./phases/phase4-mastery/exercises/03-testing-benchmarks/
```

## Hints
<details><summary>Hint 1</summary>Builder: `var b strings.Builder; for i := 0; i < n; i++ { b.WriteString(s) }; return b.String()`</details>
<details><summary>Hint 2</summary>Repeat: `return strings.Repeat(s, n)`</details>

## References
- [pkg.go.dev — testing.B](https://pkg.go.dev/testing#B)
- [Go Blog — Profiling Go programs](https://go.dev/blog/profiling-go-programs)
