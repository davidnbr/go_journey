# Project 02: Real-time Log Processor

## Overview
A channel-based log processing pipeline: parse → filter → fan-out → count.

## Run Tests
```bash
go test ./phases/phase3-concurrency/projects/02-log-processor/
```

## Skills Applied
- Channel pipelines (exercise 07 pattern)
- Fan-out to level-specific channels (exercise 06 pattern)
- `bufio.Scanner` for line-by-line reading
- `strings.SplitN` for parsing "LEVEL: message" format

## Implementation Order
1. `Parse` — scan lines, split on ": ", send LogEntry
2. `FilterLevel` — single pipeline stage
3. `CountByLevel` — drain channel into map
4. `FanOutLevels` — map of output channels + routing goroutine
