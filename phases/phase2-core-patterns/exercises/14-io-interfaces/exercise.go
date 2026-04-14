package exercise_14

import (
	"io"
	"strings"
)

// ReadAll reads all bytes from r and returns them as a string.
// TODO: use io.ReadAll(r), convert []byte to string
func ReadAll(r io.Reader) (string, error) {
	return "", nil
}

// WriteString writes s to w and returns any error.
// TODO: use io.WriteString(w, s)
func WriteString(w io.Writer, s string) error {
	return nil
}

// UpperReader returns a new io.Reader that uppercases everything read from r.
// TODO: read all from r, uppercase, return strings.NewReader of result
// Note: this is a simplified (non-streaming) implementation.
func UpperReader(r io.Reader) (io.Reader, error) {
	return nil, nil
}

var _, _ = io.ReadAll, strings.NewReader
