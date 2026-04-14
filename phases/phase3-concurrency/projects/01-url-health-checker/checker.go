package urlchecker

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Result holds the health check result for one URL.
type Result struct {
	URL        string
	StatusCode int
	Latency    time.Duration
	Error      error
}

// Check performs an HTTP GET to url with the given timeout.
// Returns a Result with status code and latency, or an error.
// TODO: create http.Client with timeout, GET url, record latency
func Check(ctx context.Context, url string, timeout time.Duration) Result {
	return Result{URL: url, Error: fmt.Errorf("not implemented")}
}

// CheckAll checks all URLs concurrently with up to workers goroutines.
// Respects ctx cancellation.
// Returns results in the same order as urls.
// TODO: worker pool pattern + context propagation
func CheckAll(ctx context.Context, urls []string, timeout time.Duration, workers int) []Result {
	results := make([]Result, len(urls))
	for i, url := range urls {
		results[i] = Result{URL: url, Error: fmt.Errorf("not implemented")}
	}
	return results
}

var _, _ = sync.WaitGroup{}, http.Get
