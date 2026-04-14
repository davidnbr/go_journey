# Project 02: Word Frequency Counter

## Overview
Build a package that reads text from any `io.Reader`, counts word frequencies, and returns the top N words.

## Learning Goals
- Use `bufio.Scanner` with `bufio.ScanWords` for token-based reading
- Build frequency maps from streaming input
- Sort structs by multiple criteria with `sort.Slice`
- Design functions around `io.Reader` for testability

## Run Tests
```bash
go test ./phases/phase1-foundations/projects/02-word-counter/
```

## Implementation Order
1. `CountWords` — scan words, lowercase, count in a map
2. `TopN` — convert map to slice, sort (desc count, asc word), truncate

## Key APIs
```go
scanner := bufio.NewScanner(r)
scanner.Split(bufio.ScanWords)
for scanner.Scan() {
    word := strings.ToLower(scanner.Text())
    // ...
}

sort.Slice(words, func(i, j int) bool {
    if words[i].Count != words[j].Count {
        return words[i].Count > words[j].Count
    }
    return words[i].Word < words[j].Word
})
```

## Skills Applied
- `io.Reader` interface, `bufio.Scanner`
- Maps, slices, `sort.Slice`
- Struct definitions and sorting by multiple keys

## Extension Ideas
- Add a `main` package that reads from `os.Stdin` or a file path argument
- Filter stop words (the, a, an, is, ...)
- Output as a histogram in the terminal
