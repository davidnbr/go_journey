# Exercise 05: Interfaces — Polymorphism

## Learning Objectives
- Implement an interface with multiple concrete types
- Write functions that operate on the interface (polymorphic dispatch)
- Return the first error from a slice of operations

## Why This Matters (DevOps context)
Alert routing in PagerDuty, OpsGenie, Slack, and email — all behind a single `Notifier` interface. This is the exact pattern used in Prometheus alertmanager receivers.

## Concept
Polymorphism via interfaces lets you add new notifier types without changing `BroadcastAll`. Any type that implements `Notify(string) error` can be passed — compile-time checked.

## Your Task
1. Implement `EmailNotifier.Notify` — just return nil (simulation)
2. Implement `SlackNotifier.Notify` — just return nil (simulation)
3. Implement `BroadcastAll` — call each notifier, return first error

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/05-interfaces-polymorphism/
```

## Hints
<details><summary>Hint 1</summary>Both notifiers just need `return nil` — they're simulated.</details>
<details><summary>Hint 2</summary>BroadcastAll: `if err := n.Notify(message); err != nil { return err }`</details>
<details><summary>Hint 3</summary>Range over `[]Notifier` — Go will dispatch to the correct concrete method.</details>

## References
- [A Tour of Go — Interfaces are implemented implicitly](https://go.dev/tour/methods/10)
