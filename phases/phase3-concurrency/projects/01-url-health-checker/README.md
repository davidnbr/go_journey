# Project 01: Concurrent URL Health Checker

## Overview
Check N URLs concurrently with a worker pool, context cancellation, and timeout per request.

## Run Tests
```bash
go test ./phases/phase3-concurrency/projects/01-url-health-checker/
```

## Skills Applied
- `net/http` client with timeout
- Worker pool (exercise 05 pattern)
- `context.Context` for cancellation + per-request timeout
- `httptest` for testing without real network

## Implementation Order
1. `Check` — single URL with timeout via `http.Client{Timeout: timeout}`
2. `CheckAll` — worker pool dispatching URLs, collecting ordered results
