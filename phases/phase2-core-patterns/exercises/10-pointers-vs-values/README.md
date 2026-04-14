# Exercise 10: Pointers vs Values

## Learning Objectives
- Understand when to pass by value vs by pointer
- Return modified copies (functional style) vs mutate in-place
- Clone structs by dereferencing

## Why This Matters (DevOps context)
Go's functional options pattern (used in AWS SDK v2, `google/options`) returns modified copies. In-place mutation is used for connection pools and caches. Knowing which pattern to use prevents subtle state-sharing bugs.

## Concept
Pass by value: caller is protected from mutation, but copies are made. Pass by pointer: efficient for large structs, required for mutation. `*ptr` dereferences to get the struct value — assigning it creates an independent copy.

## Your Task
1. `DefaultConfig` — return a Config with preset values
2. `WithDebug` — return a modified copy (do NOT change the input)
3. `SetWorkers` — mutate via pointer
4. `Clone` — return a value copy by dereferencing the pointer

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/10-pointers-vs-values/
```

## Hints
<details><summary>Hint 1</summary>WithDebug receives a copy already — just set c.Debug = true and return c.</details>
<details><summary>Hint 2</summary>Clone: `return *c` — dereferencing a pointer gives the struct value.</details>

## References
- [Effective Go — Pointers vs Values](https://go.dev/doc/effective_go#pointers_vs_values)
- [Go FAQ — Should I define methods on values or pointers?](https://go.dev/doc/faq#methods_on_values_or_pointers)
