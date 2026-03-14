# Exercise 09: sync.Mutex

## Learning Objectives
- Use `sync.Mutex` to protect shared mutable state
- Use `sync.RWMutex` for read-heavy workloads
- Understand why `go test -race` catches mutex violations

## Why This Matters (DevOps context)
Any Go service with shared state (connection pools, caches, counters, registries) needs mutex protection. The race detector (`-race`) is mandatory in CI — it catches mutex violations at runtime.

## Concept
`mu.Lock()` / `mu.Unlock()` provide exclusive access. `mu.RLock()` / `mu.RUnlock()` allow concurrent reads but block writes. Use `defer mu.Unlock()` to ensure unlock even on panic.

## Your Task
1. `SafeCounter.Inc` — add mutex locking
2. `SafeCounter.Value` — add mutex locking
3. `NewSafeMap` — initialize the inner map
4. `SafeMap.Set` and `Get` — add RWMutex locking

## Run Tests
```bash
go test -race ./phases/phase3-concurrency/exercises/09-sync-mutex/
```

## Hints
<details><summary>Hint 1</summary>`c.mu.Lock(); defer c.mu.Unlock(); c.count++`</details>
<details><summary>Hint 2</summary>NewSafeMap: `return &SafeMap{m: make(map[string]int)}`</details>

## References
- [pkg.go.dev — sync.Mutex](https://pkg.go.dev/sync#Mutex)
- [Go Blog — Share Memory By Communicating](https://go.dev/blog/share-memory-by-communicating)
