package exercise_07

// Platform returns the target OS/arch this binary was compiled for.
// The real implementation uses build tags per platform.
// For this exercise, return the runtime detection.
// TODO: use runtime.GOOS and runtime.GOARCH
func Platform() string {
	return "" // stub
}

// IsProduction reports whether the binary was built with the "production" build tag.
// Without the tag, returns false.
// The "production" build tag version returns true.
// TODO: return false (default) — the production version is in platform_prod.go
func IsProduction() bool {
	return false
}
