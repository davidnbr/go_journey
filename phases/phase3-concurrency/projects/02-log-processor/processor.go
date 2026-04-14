package logprocessor

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"sync"
)

// LogEntry is a parsed log line.
type LogEntry struct {
	Level   string
	Message string
	Raw     string
}

// Parse reads lines from r and sends LogEntry values on the returned channel.
// Closes the channel when r is exhausted or returns an error.
// TODO: goroutine + bufio.Scanner, parse "LEVEL: message" format
func Parse(r io.Reader) <-chan LogEntry {
	ch := make(chan LogEntry)
	close(ch) // stub
	return ch
}

// FilterLevel returns a new channel with only entries matching level.
// TODO: goroutine, range entries, filter by entry.Level == level
func FilterLevel(entries <-chan LogEntry, level string) <-chan LogEntry {
	out := make(chan LogEntry)
	close(out) // stub
	return out
}

// FanOutLevels fans out entries to separate channels per level.
// Returns a map of level → channel.
// levels is the set of levels to route; "other" catches the rest.
// TODO: create map of channels, goroutine routes each entry
func FanOutLevels(entries <-chan LogEntry, levels []string) map[string]<-chan LogEntry {
	return nil
}

// CountByLevel reads from entries and returns a map of level → count.
// TODO: range entries, increment map[entry.Level]
func CountByLevel(entries <-chan LogEntry) map[string]int {
	return nil
}

var _, _ = bufio.NewScanner, strings.SplitN
var _ = sync.WaitGroup{}
var _ = fmt.Sscanf
