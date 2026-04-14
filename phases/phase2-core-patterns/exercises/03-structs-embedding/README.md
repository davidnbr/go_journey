# Exercise 03: Structs — Embedding (Composition)

## Learning Objectives
- Use struct embedding to compose types
- Access promoted fields and methods
- Understand Go's "composition over inheritance" design

## Why This Matters (DevOps context)
`http.Server` embeds configuration structs. Terraform providers embed base client structs. The Logger-into-Service pattern is used throughout Go microservices for contextual logging.

## Concept
Embedding `type Dog struct { Animal; Breed string }` promotes all `Animal` fields and methods to `Dog`. There is no inheritance — `Dog` is not an `Animal` subtype. You can override promoted methods by defining them on `Dog`.

## Your Task
1. `Animal.Speak` — return `"<Name> says ..."` (any message containing the name)
2. `Dog.Fetch` — return a string containing the dog's name
3. `Logger.Log` — format `"[Prefix] message"`
4. Verify that `Service` promotes `Logger.Log` (no code to write — just understand)

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/03-structs-embedding/
```

## Hints
<details><summary>Hint 1</summary>`fmt.Sprintf("%s says ...", a.Name)` for Speak.</details>
<details><summary>Hint 2</summary>`fmt.Sprintf("[%s] %s", l.Prefix, message)` for Log.</details>
<details><summary>Hint 3</summary>Access `d.Name` directly — it's promoted from the embedded Animal.</details>

## References
- [Effective Go — Embedding](https://go.dev/doc/effective_go#embedding)
- [A Tour of Go — Methods and Interfaces](https://go.dev/tour/methods/1)
