package exercise_17

import (
	"fmt"
	"io"
	"strings"
)

// Pipeline chains readers: reads from src, uppercases, and writes to dst.
// Returns number of bytes written and any error.
// TODO: io.ReadAll from src, strings.ToUpper, io.WriteString to dst
func Pipeline(src io.Reader, dst io.Writer) (int, error) {
	return 0, fmt.Errorf("not implemented")
}

// SafeDivide divides a by b. Returns (0, ErrDivByZero) if b == 0.
// Wraps ErrDivByZero with caller context.
var ErrDivByZero = fmt.Errorf("division by zero")

func SafeDivide(a, b int) (int, error) {
	return 0, nil
}

// Config demonstrates functional options pattern.
type ServerConfig struct {
	Host    string
	Port    int
	Timeout int
}

type Option func(*ServerConfig)

func WithHost(host string) Option {
	return nil
}

func WithPort(port int) Option {
	return nil
}

// NewServerConfig applies all options to a default config.
// Default: Host="localhost", Port=8080, Timeout=30
// TODO: apply each option to the config
func NewServerConfig(opts ...Option) ServerConfig {
	return ServerConfig{}
}

var _, _ = strings.ToUpper, io.WriteString
