# Phase 2 — Core Patterns

**Bloom level**: Apply → Analyze
**Estimated time**: 20–25 hours
**Prerequisite**: Phase 1 complete

Build Go's object model: structs, methods, interfaces, pointers, error wrapping, defer/panic/recover, packages, and I/O interfaces.

## Exercises

| # | Exercise | Concept |
|---|----------|---------|
| 01 | [structs-basics](exercises/01-structs-basics/) | Struct definition, field access, initialization |
| 02 | [structs-methods](exercises/02-structs-methods/) | Value vs pointer receivers |
| 03 | [structs-embedding](exercises/03-structs-embedding/) | Composition over inheritance |
| 04 | [interfaces-basics](exercises/04-interfaces-basics/) | Interface definition, implicit satisfaction |
| 05 | [interfaces-polymorphism](exercises/05-interfaces-polymorphism/) | Polymorphic dispatch |
| 06 | [type-assertions](exercises/06-type-assertions/) | Type assertions, type switches |
| 07 | [errors-custom](exercises/07-errors-custom/) | Custom error types, sentinel errors |
| 08 | [errors-wrapping](exercises/08-errors-wrapping/) | `%w`, `errors.Is`, `errors.As` |
| 09 | [pointers-basics](exercises/09-pointers-basics/) | `&`, `*`, nil pointers |
| 10 | [pointers-vs-values](exercises/10-pointers-vs-values/) | When to use pointer vs value receivers |
| 11 | [defer-panic-recover](exercises/11-defer-panic-recover/) | Defer ordering, panic/recover pattern |
| 12 | [packages-design](exercises/12-packages-design/) | Visibility, `init()`, package naming |
| 13 | [modules-management](exercises/13-modules-management/) | `go.mod`, versioning, `go get` |
| 14 | [io-interfaces](exercises/14-io-interfaces/) | `io.Reader`, `io.Writer`, composing I/O |
| 15 | [goroutines-preview](exercises/15-goroutines-preview/) | Conceptual goroutine intro |
| 16 | [review-2a](exercises/16-review-2a/) | Mix of exercises 01–08 |
| 17 | [review-2b](exercises/17-review-2b/) | Mix of exercises 09–15 |
| 18 | [review-2-hard](exercises/18-review-2-hard/) | Interleaved, application level |

## Projects

| # | Project | Skills |
|---|---------|--------|
| 01 | [Todo CLI](projects/01-todo-cli/) | Structs, interfaces, JSON persistence |
| 02 | [Simple HTTP Server](projects/02-http-server/) | `net/http`, JSON, middleware concept |
| 03 | [Mini Bank System](projects/03-mini-bank/) | Interfaces, multiple packages, custom errors |
