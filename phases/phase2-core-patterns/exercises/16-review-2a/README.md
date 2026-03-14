# Exercise 16: Review 2a — Structs, Methods, Interfaces, Errors

## Learning Objectives
Reinforce Phase 2 concepts 01–08 by building a generic stack data structure.

## Concepts Covered
- Struct with unexported fields (`items`)
- Pointer receiver methods
- Custom sentinel error
- Slice manipulation (append, shrink)

## Your Task
Implement a LIFO stack: Push, Pop, Peek, Len.

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/16-review-2a/
```

## Hints
<details><summary>Hint 1</summary>Pop: `n := len(s.items); item := s.items[n-1]; s.items = s.items[:n-1]; return item, nil`</details>
<details><summary>Hint 2</summary>Peek: same as Pop but don't shrink the slice.</details>
