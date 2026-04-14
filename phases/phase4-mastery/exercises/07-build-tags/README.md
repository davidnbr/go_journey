# Exercise 07: Build Tags and Cross-Compilation

## Learning Objectives
- Use build constraints (`//go:build`) to include/exclude files
- Use `runtime.GOOS` and `runtime.GOARCH` for runtime platform detection
- Understand cross-compilation: `GOOS=linux GOARCH=amd64 go build`

## Why This Matters (DevOps context)
CI pipelines build binaries for multiple platforms. Build tags enable platform-specific implementations, feature flags, and debug vs production code paths — used in Docker, Kubernetes, and Terraform.

## Concept
`//go:build linux` at the top of a file includes it only on Linux. `go build -tags production` enables files tagged `//go:build production`. `runtime.GOOS` detects the OS at runtime.

## Your Task
1. Implement `Platform()` using `runtime.GOOS + "/" + runtime.GOARCH`
2. `IsProduction()` returns false by default — this is correct (no production tag)

## Cross-Compile Example
```bash
GOOS=darwin GOARCH=arm64 go build ./phases/phase4-mastery/exercises/07-build-tags/
GOOS=windows GOARCH=amd64 go build ./phases/phase4-mastery/exercises/07-build-tags/
```

## Run Tests
```bash
go test ./phases/phase4-mastery/exercises/07-build-tags/
```

## Hints
<details><summary>Hint 1</summary>`import "runtime"; return runtime.GOOS + "/" + runtime.GOARCH`</details>

## References
- [Go Reference — Build constraints](https://pkg.go.dev/cmd/go#hdr-Build_constraints)
- [Go Blog — Cross-compilation](https://go.dev/wiki/GoArm)
