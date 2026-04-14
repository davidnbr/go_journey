package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// Operation represents a supported arithmetic operation.
type Operation string

const (
	OpAdd      Operation = "add"
	OpSubtract Operation = "subtract"
	OpMultiply Operation = "multiply"
	OpDivide   Operation = "divide"
	OpSqrt     Operation = "sqrt"
)

// calculate performs the given operation on a and b.
// For sqrt, only a is used.
// Returns the result and an error.
// TODO: implement each case using a switch on op
func calculate(op Operation, a, b float64) (float64, error) {
	return 0, fmt.Errorf("not implemented")
}

// parseArgs parses os.Args and returns (op, a, b, error).
// Usage:
//
//	add|subtract|multiply|divide <a> <b>
//	sqrt <a>
//
// TODO: validate arg count, parse floats with strconv.ParseFloat
func parseArgs(args []string) (Operation, float64, float64, error) {
	return "", 0, 0, fmt.Errorf("not implemented")
}

func main() {
	op, a, b, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\nUsage: calc <op> <a> [b]\n  ops: add subtract multiply divide sqrt\n", err)
		os.Exit(1)
	}
	result, err := calculate(op, a, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%.6g\n", result)
}

// Keep imports used by the student's implementation
var _, _ = strconv.ParseFloat, math.Sqrt
