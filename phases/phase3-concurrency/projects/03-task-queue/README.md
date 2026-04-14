# Project 03: Mini Task Queue

## Overview
A concurrent task queue with bounded workers, submit, and graceful shutdown.

## Run Tests
```bash
go test ./phases/phase3-concurrency/projects/03-task-queue/
```

## Skills Applied
- Goroutines + channels for the worker loop
- `sync.WaitGroup` for graceful shutdown
- `sync.Once` to close results channel exactly once
- Buffered job channel for backpressure
- Context propagation to tasks

## Implementation Order
1. `NewQueue` — init channels and start workers
2. `Start` — launch N worker goroutines reading from q.jobs
3. `Submit` — send to q.jobs, handle full buffer
4. `Stop` — close q.jobs, wait, close q.results
