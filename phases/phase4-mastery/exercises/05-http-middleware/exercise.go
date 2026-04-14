package exercise_05

import (
	"fmt"
	"net/http"
	"time"
)

// Middleware wraps an http.Handler.
type Middleware func(http.Handler) http.Handler

// Chain applies middlewares in order: first middleware is outermost.
// TODO: apply each middleware wrapping the next, right to left
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	return h // stub: no middleware applied
}

// Logger logs method, path, and duration for each request.
// TODO: record start time, call next, log elapsed
func Logger(next http.Handler) http.Handler {
	return next // stub: pass-through
}

// Recovery recovers from panics and returns 500.
// TODO: defer recover(), if panic write 500
func Recovery(next http.Handler) http.Handler {
	return next // stub: pass-through
}

// Auth checks for "Authorization: Bearer secret" header.
// Returns 401 if missing or wrong.
// TODO: check header, return 401, else call next
func Auth(secret string) Middleware {
	return func(next http.Handler) http.Handler {
		return next // stub: pass-through
	}
}

var _ = fmt.Sprintf
var _ = time.Now
