# Project 02: DevOps CLI Tool — URL Health Checker

## Overview
A production-ready CLI tool that checks multiple HTTP endpoints concurrently and exits with a non-zero code if any are unhealthy.

## Run Tests
```bash
go test ./phases/phase4-mastery/projects/02-devops-cli/
```

## Skills Applied
- `net/http` client with configurable timeout
- Worker pool for concurrent checks
- Context cancellation propagation
- Exit code semantics (`os.Exit(1)` on failure)
- `httptest` for test isolation

## Implementation Order
1. `CheckURL` — single URL, measure latency, return Result
2. `RunChecks` — worker pool, collect in order
3. `ExitCode` — 0 if all status 2xx, else 1
4. `PrintResults` — already works, improve formatting as extension

## Extension Ideas
- Add `--json` flag for machine-readable output
- Add retry logic with exponential backoff
- Read URLs from a YAML config file
- Add Prometheus metrics endpoint
