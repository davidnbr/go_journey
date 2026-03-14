package exercise_07

import (
	"runtime"
	"strings"
	"testing"
)

func TestPlatform(t *testing.T) {
	got := Platform()
	if got == "" {
		t.Error("Platform() should not be empty")
	}
	// Should contain GOOS/GOARCH info
	if !strings.Contains(got, runtime.GOOS) {
		t.Errorf("Platform() = %q, should contain %q", got, runtime.GOOS)
	}
}

func TestIsProduction(t *testing.T) {
	// Without -tags production, should return false
	if IsProduction() {
		t.Error("IsProduction() without build tag should be false")
	}
}
