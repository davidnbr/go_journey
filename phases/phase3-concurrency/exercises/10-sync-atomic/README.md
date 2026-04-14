# Exercise 10: sync/atomic

## Learning Objectives
- Use `sync/atomic` for lock-free integer operations
- Understand when atomic operations are preferable to mutexes
- Implement compare-and-swap (CAS) for optimistic updates

## Why This Matters (DevOps context)
Atomic counters are used for metrics (request counts, error counts) where mutex overhead is undesirable. CAS underpins lock-free data structures and optimistic concurrency control.

## Concept
`atomic.AddInt64(&v, delta)`, `atomic.LoadInt64(&v)`, `atomic.StoreInt64(&v, val)`, `atomic.CompareAndSwapInt64(&v, old, new)` — all operate without locks, using CPU-level atomics.

## Your Task
Implement all four `AtomicCounter` methods using the `sync/atomic` package.

## Run Tests
```bash
go test -race ./phases/phase3-concurrency/exercises/10-sync-atomic/
```

## Hints
<details><summary>Hint 1</summary>`atomic.AddInt64(&c.value, 1)` for Inc; `atomic.LoadInt64(&c.value)` for Value.</details>
<details><summary>Hint 2</summary>`atomic.CompareAndSwapInt64(&c.value, oldVal, newVal)`</details>

## References
- [pkg.go.dev — sync/atomic](https://pkg.go.dev/sync/atomic)
