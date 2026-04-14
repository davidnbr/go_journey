package exercise_12

// Exported and unexported identifiers demo.

// exported — visible outside the package
const Version = "0.0.0"

// unexported — only visible within this package
const internalBuildID = "dev"

// Config holds settings.
type Config struct {
	Host   string // exported
	Port   int    // exported
	secret string
}

// NewConfig returns a Config with defaults and sets the unexported secret.
// TODO: return Config{Host: "localhost", Port: 8080, secret: "changeme"}
func NewConfig() Config {
	return Config{}
}

// IsValid returns true if Host is non-empty and Port > 0.
// TODO: implement
func (c Config) IsValid() bool {
	return false
}

// GetVersion returns the Version constant.
// TODO: return Version
func GetVersion() string {
	return ""
}
