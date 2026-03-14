# Exercise 13: Modules Management

## Learning Objectives
- Understand `go.mod` structure: module path, go version, dependencies
- Know how `go get`, `go mod tidy`, and module versioning work
- Implement simple semantic version detection

## Why This Matters (DevOps context)
Go modules are how you manage dependencies in every Go project. Understanding module paths, versions, and `go.sum` is essential for reproducible builds in CI/CD pipelines.

## Concept
`go.mod` declares the module path (used in import paths), the minimum Go version, and external dependencies. `go mod tidy` removes unused and adds missing dependencies. Semantic versioning: `v<major>.<minor>.<patch>`.

## Your Task
1. `ModuleInfo` — return this module's path
2. `GoVersion` — return the Go version from go.mod
3. `IsSemanticVersion` — validate the semver format

## Run Tests
```bash
go test ./phases/phase2-core-patterns/exercises/13-modules-management/
```

## Hints
<details><summary>Hint 1</summary>Module path is in `/home/dbecerra/Documents/Projects/Golang/go_journey/go.mod` — look it up.</details>
<details><summary>Hint 2</summary>IsSemanticVersion: `strings.HasPrefix(s, "v") && strings.Count(s, ".") >= 2`</details>

## References
- [Go Modules Reference](https://go.dev/ref/mod)
- [Go Blog — Using Go Modules](https://go.dev/blog/using-go-modules)
