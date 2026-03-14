package exercise_08

import (
	"errors"
	"fmt"
)

// ErrPermission is a sentinel for permission denied.
var ErrPermission = errors.New("permission denied")

// ErrNotFound is a sentinel for resource not found.
var ErrNotFound = errors.New("not found")

// OpenConfig simulates opening a config file.
// If name == "missing.yaml" → wrap ErrNotFound with context.
// If name == "secret.yaml" → wrap ErrPermission with context.
// Otherwise → return nil.
// TODO: use fmt.Errorf with %w to wrap
func OpenConfig(name string) error {
	return nil
}

// LoadConfig calls OpenConfig and wraps any error with "loadConfig: ..."
// TODO: if err != nil, return fmt.Errorf("loadConfig: %w", err)
func LoadConfig(name string) error {
	return nil
}

var _, _ = fmt.Errorf, errors.Is
