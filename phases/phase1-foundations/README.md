# Phase 1 — Foundations

**Bloom level**: Remember → Apply
**Estimated time**: 15–20 hours

Build fluency with Go's core syntax: types, control flow, functions, closures, slices, maps, strings, and error handling. Every exercise is a failing test you make pass.

## Exercises

| # | Exercise | Concept |
|---|----------|---------|
| 01 | [hello-world](exercises/01-hello-world/) | `package main`, `fmt.Println`, compilation |
| 02 | [variables](exercises/02-variables/) | `:=`, `var`, zero values, multiple assignment |
| 03 | [types](exercises/03-types/) | Basic types, type conversion |
| 04 | [constants](exercises/04-constants/) | `const`, `iota` |
| 05 | [functions-basics](exercises/05-functions-basics/) | Single + multiple return values |
| 06 | [functions-advanced](exercises/06-functions-advanced/) | Named returns, variadic, first-class functions |
| 07 | [closures](exercises/07-closures/) | Closure capture, counter pattern |
| 08 | [control-if](exercises/08-control-if/) | if/else, if-with-init |
| 09 | [control-switch](exercises/09-control-switch/) | switch, fallthrough, type switch |
| 10 | [loops](exercises/10-loops/) | for, range, break/continue |
| 11 | [arrays-slices](exercises/11-arrays-slices/) | Fixed arrays vs slices, len/cap |
| 12 | [slices-advanced](exercises/12-slices-advanced/) | append, copy, slicing internals |
| 13 | [maps](exercises/13-maps/) | CRUD, comma-ok idiom, iteration |
| 14 | [strings-runes](exercises/14-strings-runes/) | Immutability, UTF-8, rune iteration |
| 15 | [errors-basics](exercises/15-errors-basics/) | `error` interface, `errors.New`, error returns |
| 16 | [review-1](exercises/16-review-1/) | **Spaced repetition**: mix of exercises 01–15 |

## Projects

| # | Project | Skills |
|---|---------|--------|
| 01 | [CLI Calculator](projects/01-cli-calculator/) | `os.Args`, type conversion, error handling |
| 02 | [Word Frequency Counter](projects/02-word-counter/) | File I/O, maps, strings, error handling |

## How to Work Through This Phase

1. Start at exercise 01 — read the README, then attempt the stub
2. Run `make check PHASE=phase1-foundations EX=<exercise-dir>`
3. Reveal hints progressively only if stuck (3 levels in each README)
4. Do exercise 16 (review) after completing 01–15 without looking back
