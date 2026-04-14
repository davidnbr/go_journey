# Exercise 18: Review 2 — Hard (Application Level)

## Learning Objectives
Combine structs, interfaces, pointers, error types, and sync primitives in two real-world patterns.

## Concepts Covered
- EventBus: struct with map state, pointer receivers, RWMutex for thread safety
- LimitedWriter: implementing `io.Writer`, stateful byte tracking, sentinel error

## Your Task
1. `EventBus`: implement `NewEventBus`, `Subscribe`, `Publish`
2. `LimitedWriter`: implement `Write` to enforce a byte limit

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/18-review-2-hard/
```

## Hints
<details><summary>Hint 1</summary>NewEventBus: `return &EventBus{handlers: make(map[string][]func(interface{}))}`</details>
<details><summary>Hint 2</summary>Publish: use `b.mu.RLock(); h := b.handlers[event]; b.mu.RUnlock()` then call each handler.</details>
<details><summary>Hint 3</summary>LimitedWriter.Write: `if lw.written >= lw.Limit { return 0, ErrWriterFull }`; write min(remaining, len(p)) bytes.</details>
