package exercise_12

import "testing"

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version should not be empty")
	}
	if Version != "0.0.0" {
		t.Errorf("Version = %q, want 0.0.0", Version)
	}
}

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	if c.Host != "localhost" {
		t.Errorf("Host = %q, want localhost", c.Host)
	}
	if c.Port != 8080 {
		t.Errorf("Port = %d, want 8080", c.Port)
	}
	// unexported field: accessible within same package test
	if c.secret != "changeme" {
		t.Errorf("secret = %q, want changeme", c.secret)
	}
}

func TestIsValid(t *testing.T) {
	if !NewConfig().IsValid() {
		t.Error("default config should be valid")
	}
	if (Config{}).IsValid() {
		t.Error("empty config should be invalid")
	}
}

func TestGetVersion(t *testing.T) {
	if got := GetVersion(); got != "0.0.0" {
		t.Errorf("GetVersion() = %q, want 0.0.0", got)
	}
}
