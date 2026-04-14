package exercise_10

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

// Cache is a generic thread-safe cache with TTL awareness (simplified: no expiry).
type Cache[K comparable, V any] struct {
	mu    sync.RWMutex
	items map[K]V
}

// NewCache creates an empty Cache.
// TODO: initialize items map
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{}
}

// Set stores a key-value pair.
// TODO: write lock, store
func (c *Cache[K, V]) Set(key K, value V) {
	// no-op stub
}

// Get retrieves a value.
// TODO: read lock, comma-ok lookup
func (c *Cache[K, V]) Get(key K) (V, bool) {
	var zero V
	return zero, false
}

// ErrServiceUnavailable is returned when the HTTP service is down.
var ErrServiceUnavailable = errors.New("service unavailable")

// HealthChecker checks URLs with concurrency and timeout.
type HealthChecker struct {
	client  *http.Client
	workers int
}

// NewHealthChecker creates a HealthChecker.
// TODO: &HealthChecker{client: &http.Client{Timeout: timeout}, workers: workers}
func NewHealthChecker(workers int) *HealthChecker {
	return &HealthChecker{}
}

// CheckAll checks all URLs concurrently, returns map[url]error (nil=healthy).
// TODO: worker pool + context, GET each URL
func (h *HealthChecker) CheckAll(ctx context.Context, urls []string) map[string]error {
	result := make(map[string]error)
	for _, url := range urls {
		result[url] = fmt.Errorf("not implemented")
	}
	return result
}

var _ = http.Get
