package exercise_10

// Config holds application settings.
type Config struct {
	Debug   bool
	Workers int
	Host    string
}

// DefaultConfig returns a Config with sensible defaults.
// TODO: return Config{Debug: false, Workers: 4, Host: "localhost"}
func DefaultConfig() Config {
	return Config{}
}

// WithDebug returns a new Config with Debug set to true.
// Use value receiver — returns a modified copy.
// TODO: set c.Debug = true, return c
func WithDebug(c Config) Config {
	return c
}

// SetWorkers modifies the Workers field via pointer.
// Use pointer receiver — mutates in-place.
// TODO: *c.Workers = n ... actually: c.Workers = n
func SetWorkers(c *Config, n int) {
	// no-op stub
}

// Clone returns a deep copy of c.
// TODO: dereference the pointer to get a value copy: return *c
func Clone(c *Config) Config {
	return Config{}
}
