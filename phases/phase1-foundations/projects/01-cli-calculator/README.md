# Project 01: CLI Calculator

## Overview
Build a command-line calculator that reads an operation and operands from `os.Args` and prints the result.

## Learning Goals
- Parse command-line arguments with `os.Args`
- Convert strings to float64 with `strconv.ParseFloat`
- Use `switch` for operation dispatch
- Return structured errors; print to stderr and exit with non-zero code on failure

## Usage (once implemented)
```bash
go run ./phases/phase1-foundations/projects/01-cli-calculator/ add 3 4
# → 7
go run ./phases/phase1-foundations/projects/01-cli-calculator/ sqrt 16
# → 4
go run ./phases/phase1-foundations/projects/01-cli-calculator/ divide 10 0
# → Error: division by zero
```

## Run Tests
```bash
go test ./phases/phase1-foundations/projects/01-cli-calculator/
```

## Implementation Order
1. Implement `calculate` — switch over `Operation` constants
2. Implement `parseArgs` — validate argc, parse floats
3. Wire up `main` — it's already done for you

## Skills Applied
- `os.Args`, `strconv.ParseFloat`, `fmt.Fprintf`, `os.Exit`
- Error returns, sentinel errors
- Constants and switch

## Extension Ideas
- Add a `pow` operation (use `math.Pow`)
- Support piped input (`os.Stdin`) for batch calculations
- Pretty-print integers without decimal point
