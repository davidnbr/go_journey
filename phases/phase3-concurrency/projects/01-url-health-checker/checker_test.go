package urlchecker

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	result := Check(context.Background(), srv.URL, 5*time.Second)
	if result.Error != nil {
		t.Fatalf("Check error: %v", result.Error)
	}
	if result.StatusCode != http.StatusOK {
		t.Errorf("Check status = %d, want 200", result.StatusCode)
	}
	if result.Latency == 0 {
		t.Error("Check latency should be > 0")
	}
}

func TestCheckTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	result := Check(context.Background(), srv.URL, 20*time.Millisecond)
	if result.Error == nil {
		t.Error("Check with short timeout should return error")
	}
}

func TestCheckAll(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	urls := []string{srv.URL, srv.URL, srv.URL}
	results := CheckAll(context.Background(), urls, 5*time.Second, 2)

	if len(results) != 3 {
		t.Fatalf("CheckAll returned %d results, want 3", len(results))
	}
	for i, r := range results {
		if r.Error != nil {
			t.Errorf("results[%d] error: %v", i, r.Error)
		}
		if r.URL != srv.URL {
			t.Errorf("results[%d].URL = %q, want %q", i, r.URL, srv.URL)
		}
	}
}
