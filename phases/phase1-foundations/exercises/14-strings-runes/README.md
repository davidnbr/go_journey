# Exercise 14: Strings and Runes

## Learning Objectives
- Understand that Go strings are immutable byte sequences (UTF-8)
- Use `range` over strings to iterate by rune (code point), not byte
- Convert between `string` and `[]rune` for Unicode-safe manipulation

## Why This Matters (DevOps context)
Infrastructure code processes hostnames, log messages, and user-facing strings. Multi-byte UTF-8 support matters when processing international content, Kubernetes resource names, or Unicode log data.

## Concept
`len(s)` counts bytes. `range s` iterates runes (Unicode code points). A rune is an alias for `int32`. To reverse or index by character, convert to `[]rune` first.

## Your Task
1. `IsPalindrome` — case-insensitive palindrome check
2. `CountRunes` — count Unicode code points (not bytes)
3. `ReverseString` — reverse by rune, handling multi-byte chars
4. `ContainsAll` — check all substrings are present

## Run Tests
```bash
go test ./phases/phase1-foundations/exercises/14-strings-runes/
```

## Hints
<details><summary>Hint 1 — Direction</summary>
ReverseString: `runes := []rune(s)` — then reverse in-place, then `string(runes)`.
</details>

<details><summary>Hint 2 — Pseudocode</summary>

```go
runes := []rune(s)
for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
}
return string(runes)
```
</details>

<details><summary>Hint 3 — Solution approach</summary>
IsPalindrome: lowercase the string, reverse it, compare. `strings.ToLower` + `ReverseString` reuse.
</details>

## References
- [A Tour of Go — Strings and bytes](https://go.dev/tour/basics/11)
- [Go Blog — Strings, bytes, runes and characters](https://go.dev/blog/strings)
- [pkg.go.dev — strings package](https://pkg.go.dev/strings)
