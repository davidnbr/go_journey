# Phase 3 — Concurrency & Advanced

**Bloom level**: Analyze → Evaluate
**Estimated time**: 20–25 hours
**Prerequisite**: Phase 2 complete

Master Go's concurrency primitives: goroutines, channels, select, worker pools, pipelines, context cancellation, sync primitives, and an introduction to generics.

## Exercises

| # | Exercise | Concept |
|---|----------|---------|
| 01 | [goroutines-lifecycle](exercises/01-goroutines-lifecycle/) | Creation, goroutine leaks, WaitGroup |
| 02 | [channels-unbuffered](exercises/02-channels-unbuffered/) | Send/receive synchronization |
| 03 | [channels-buffered](exercises/03-channels-buffered/) | Buffer semantics, deadlock avoidance |
| 04 | [select-statement](exercises/04-select-statement/) | Multiplexing channels, default case |
| 05 | [worker-pool](exercises/05-worker-pool/) | Fixed goroutine pool pattern |
| 06 | [fan-out-fan-in](exercises/06-fan-out-fan-in/) | Distributing and merging work |
| 07 | [pipeline](exercises/07-pipeline/) | Staged data transformation |
| 08 | [context-cancellation](exercises/08-context-cancellation/) | `context.WithCancel`, `WithTimeout` |
| 09 | [sync-mutex](exercises/09-sync-mutex/) | `sync.Mutex`, `RWMutex` |
| 10 | [sync-atomic](exercises/10-sync-atomic/) | `sync/atomic` operations |
| 11 | [sync-waitgroup](exercises/11-sync-waitgroup/) | `sync.WaitGroup` patterns |
| 12 | [race-conditions](exercises/12-race-conditions/) | `-race` flag, detecting data races |
| 13 | [generics-basics](exercises/13-generics-basics/) | Type parameters, constraints (Go 1.18+) |
| 14 | [reflection-basics](exercises/14-reflection-basics/) | `reflect` package, use cases |
| 15 | [review-3](exercises/15-review-3/) | Concurrency patterns interleaved |

## Projects

| # | Project | Skills |
|---|---------|--------|
| 01 | [URL Health Checker](projects/01-url-health-checker/) | Goroutines, channels, context, HTTP |
| 02 | [Log Processor](projects/02-log-processor/) | Pipelines, fan-out, worker pools |
| 03 | [Task Queue](projects/03-task-queue/) | Goroutines, channels, sync primitives |
