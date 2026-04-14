# Exercise 01: Structs — Basics

## Learning Objectives
- Define struct types with named fields
- Initialize structs with field-name syntax
- Define methods on struct types

## Why This Matters (DevOps context)
Structs model infrastructure resources: an EC2 instance, a Kubernetes pod spec, a Terraform resource block. Named field initialization prevents field-order bugs in large structs.

## Concept
Structs are value types in Go. Methods are defined with a receiver: `func (r Rectangle) Area() float64`. Embedding structs (next exercise) gives you composition without inheritance.

## Your Task
1. `NewPoint` — return a Point with the given X, Y
2. `Area` — multiply Width by Height
3. `Contains` — check if a point falls within the rectangle bounds

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/01-structs-basics/
```

## Hints
<details><summary>Hint 1</summary>Return `Point{X: x, Y: y}`</details>
<details><summary>Hint 2</summary>For Contains: `pt.X >= r.TopLeft.X && pt.X <= r.TopLeft.X+r.Width`</details>
<details><summary>Hint 3</summary>Both X and Y conditions must hold — combine with `&&`.</details>

## References
- [A Tour of Go — Structs](https://go.dev/tour/moretypes/2)
- [Effective Go — Structs](https://go.dev/doc/effective_go#composite_literals)
