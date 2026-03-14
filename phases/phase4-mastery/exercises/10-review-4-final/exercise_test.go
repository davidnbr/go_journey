package exercise_10

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewCache[string, int]()
	if c == nil {
		t.Fatal("NewCache returned nil")
	}
	c.Set("a", 1)
	c.Set("b", 2)

	v, ok := c.Get("a")
	if !ok || v != 1 {
		t.Errorf("Get(a) = (%d, %v), want (1, true)", v, ok)
	}
	_, ok = c.Get("missing")
	if ok {
		t.Error("Get(missing) should return ok=false")
	}
}

func TestCacheConcurrent(t *testing.T) {
	c := NewCache[int, string]()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Set(i, "val")
			c.Get(i)
		}()
	}
	wg.Wait()
}

func TestHealthCheckerAll(t *testing.T) {
	healthy := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer healthy.Close()

	checker := NewHealthChecker(3)
	results := checker.CheckAll(context.Background(), []string{healthy.URL, healthy.URL})

	for url, err := range results {
		if err != nil {
			t.Errorf("CheckAll(%s) unexpected error: %v", url, err)
		}
	}
}
