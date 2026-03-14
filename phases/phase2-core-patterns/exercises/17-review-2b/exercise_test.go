package exercise_17

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestPipeline(t *testing.T) {
	src := strings.NewReader("hello world")
	var dst bytes.Buffer
	n, err := Pipeline(src, &dst)
	if err != nil {
		t.Fatalf("Pipeline error: %v", err)
	}
	if dst.String() != "HELLO WORLD" {
		t.Errorf("Pipeline output = %q, want HELLO WORLD", dst.String())
	}
	if n <= 0 {
		t.Errorf("Pipeline wrote %d bytes, want > 0", n)
	}
}

func TestSafeDivide(t *testing.T) {
	v, err := SafeDivide(10, 2)
	if err != nil || v != 5 {
		t.Errorf("SafeDivide(10,2) = (%d, %v), want (5, nil)", v, err)
	}
	_, err = SafeDivide(5, 0)
	if !errors.Is(err, ErrDivByZero) {
		t.Errorf("SafeDivide(5,0) should wrap ErrDivByZero, got %v", err)
	}
}

func TestFunctionalOptions(t *testing.T) {
	cfg := NewServerConfig(WithHost("prod.example.com"), WithPort(9090))
	if cfg.Host != "prod.example.com" {
		t.Errorf("Host = %q, want prod.example.com", cfg.Host)
	}
	if cfg.Port != 9090 {
		t.Errorf("Port = %d, want 9090", cfg.Port)
	}
	if cfg.Timeout != 30 {
		t.Errorf("default Timeout = %d, want 30", cfg.Timeout)
	}
}
