# Exercise 14: I/O Interfaces

## Learning Objectives
- Use `io.Reader` and `io.Writer` as universal I/O abstractions
- Read from any reader with `io.ReadAll`
- Write to any writer with `io.WriteString`
- Compose readers/writers

## Why This Matters (DevOps context)
`io.Reader`/`io.Writer` are the foundation of Go I/O: HTTP request/response bodies, file handles, network connections, compression streams, and test buffers all implement them. Writing against these interfaces makes code testable without disk or network.

## Concept
`io.Reader` has one method: `Read(p []byte) (n int, err error)`. `io.Writer` has `Write(p []byte) (n int, err error)`. `bytes.Buffer`, `strings.NewReader`, `os.File`, and `net.Conn` all implement these.

## Your Task
1. `ReadAll` — drain a reader to a string
2. `WriteString` — write a string to a writer
3. `UpperReader` — transform: read all, uppercase, wrap in new reader

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/14-io-interfaces/
```

## Hints
<details><summary>Hint 1</summary>`data, err := io.ReadAll(r); return string(data), err`</details>
<details><summary>Hint 2</summary>`_, err := io.WriteString(w, s); return err`</details>
<details><summary>Hint 3</summary>UpperReader: read all, `strings.ToUpper(content)`, `return strings.NewReader(upper), nil`</details>

## References
- [pkg.go.dev — io package](https://pkg.go.dev/io)
- [Go Blog — The Go I/O model](https://go.dev/blog/io2013)
