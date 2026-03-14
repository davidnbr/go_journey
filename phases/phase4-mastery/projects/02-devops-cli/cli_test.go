package devcli

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheckURL(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	result := CheckURL(context.Background(), srv.URL, 5*time.Second)
	if result.Error != "" {
		t.Errorf("CheckURL error: %v", result.Error)
	}
	if result.Status != http.StatusOK {
		t.Errorf("CheckURL status = %d, want 200", result.Status)
	}
	if result.Latency == 0 {
		t.Error("CheckURL latency should be > 0")
	}
}

func TestRunChecks(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer srv.Close()

	cfg := Config{
		URLs:    []string{srv.URL, srv.URL, srv.URL},
		Timeout: 5 * time.Second,
		Workers: 2,
	}
	results := RunChecks(context.Background(), cfg)
	if len(results) != 3 {
		t.Fatalf("RunChecks returned %d results, want 3", len(results))
	}
	for i, r := range results {
		if r.Error != "" {
			t.Errorf("results[%d] error: %v", i, r.Error)
		}
	}
}

func TestExitCode(t *testing.T) {
	healthy := []Result{{URL: "a", Status: 200}, {URL: "b", Status: 204}}
	if code := ExitCode(healthy); code != 0 {
		t.Errorf("ExitCode(all healthy) = %d, want 0", code)
	}
	withError := []Result{{URL: "a", Status: 200}, {URL: "b", Error: "timeout"}}
	if code := ExitCode(withError); code != 1 {
		t.Errorf("ExitCode(with error) = %d, want 1", code)
	}
}
