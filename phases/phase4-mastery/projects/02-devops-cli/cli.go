// Package devcli provides a DevOps CLI tool skeleton.
// It checks HTTP endpoint health and reports results.
package devcli

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

// Config holds CLI configuration.
type Config struct {
	URLs    []string
	Timeout time.Duration
	Workers int
	Output  string // "text" or "json"
}

// DefaultConfig returns sensible defaults.
func DefaultConfig() Config {
	return Config{
		Timeout: 5 * time.Second,
		Workers: 5,
		Output:  "text",
	}
}

// Result is the outcome of a single health check.
type Result struct {
	URL     string
	Status  int
	Latency time.Duration
	Error   string
}

// CheckURL performs a single HTTP GET and returns a Result.
// TODO: use http.Client with timeout, measure latency
func CheckURL(ctx context.Context, url string, timeout time.Duration) Result {
	return Result{URL: url, Error: "not implemented"}
}

// RunChecks checks all URLs in cfg concurrently.
// Returns slice of Results in same order as cfg.URLs.
// TODO: worker pool, collect results in order
func RunChecks(ctx context.Context, cfg Config) []Result {
	results := make([]Result, len(cfg.URLs))
	for i, url := range cfg.URLs {
		results[i] = Result{URL: url, Error: "not implemented"}
	}
	return results
}

// PrintResults prints results to stdout.
// TODO: for each result, print URL + status or error
func PrintResults(results []Result) {
	for _, r := range results {
		if r.Error != "" {
			fmt.Printf("FAIL %s: %s\n", r.URL, r.Error)
		} else {
			fmt.Printf("OK   %s: %d (%v)\n", r.URL, r.Status, r.Latency)
		}
	}
}

// ExitCode returns 0 if all results are healthy (status 200-299), 1 otherwise.
// TODO: check each result
func ExitCode(results []Result) int {
	return 1 // stub: always failure
}

var _ = os.Exit
var _ = http.Get
