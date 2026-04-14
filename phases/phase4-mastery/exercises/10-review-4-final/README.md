# Exercise 10: Review 4 — Final (All Phases)

## Learning Objectives
Demonstrate mastery across all four phases: generics, sync primitives, HTTP, concurrency, and testing patterns combined in one exercise.

## Concepts Covered
- Generic `Cache[K, V]` with `sync.RWMutex`
- `HealthChecker` with worker pool, `net/http` client, context
- Table-driven test design throughout all phases

## Your Task
1. `NewCache` — initialize with generic map
2. `Cache.Set` / `Get` — RWMutex-protected read/write
3. `NewHealthChecker` — http.Client with timeout
4. `CheckAll` — concurrent health checks, collect results

## Run Tests
```bash
go test -race ./phases/phase4-mastery/exercises/10-review-4-final/
```

## Hints
<details><summary>Hint 1</summary>Cache init: `return &Cache[K, V]{items: make(map[K]V)}`</details>
<details><summary>Hint 2</summary>CheckAll: use a mutex-protected map for results, WaitGroup for goroutines, channel for bounded workers.</details>
