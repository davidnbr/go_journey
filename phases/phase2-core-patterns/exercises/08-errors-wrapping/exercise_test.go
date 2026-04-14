package exercise_08

import (
	"errors"
	"testing"
)

func TestOpenConfig(t *testing.T) {
	if err := OpenConfig("app.yaml"); err != nil {
		t.Errorf("OpenConfig(app.yaml) unexpected error: %v", err)
	}
	err := OpenConfig("missing.yaml")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("OpenConfig(missing.yaml) should wrap ErrNotFound, got: %v", err)
	}
	err = OpenConfig("secret.yaml")
	if !errors.Is(err, ErrPermission) {
		t.Errorf("OpenConfig(secret.yaml) should wrap ErrPermission, got: %v", err)
	}
}

func TestLoadConfig(t *testing.T) {
	err := LoadConfig("missing.yaml")
	if err == nil {
		t.Fatal("LoadConfig(missing.yaml) should return error")
	}
	// errors.Is traverses the chain
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("LoadConfig error chain should contain ErrNotFound, got: %v", err)
	}
	// Error message should include both wrapper and wrapped context
	msg := err.Error()
	if msg == "" {
		t.Error("error message should not be empty")
	}
}
