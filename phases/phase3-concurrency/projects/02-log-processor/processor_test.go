package logprocessor

import (
	"strings"
	"testing"
)

const testLog = `ERROR: disk full
INFO: service started
WARN: memory at 80%
ERROR: connection refused
INFO: request processed
DEBUG: cache hit`

func TestParse(t *testing.T) {
	entries := Parse(strings.NewReader(testLog))
	var got []LogEntry
	for e := range entries {
		got = append(got, e)
	}
	if len(got) != 6 {
		t.Fatalf("Parse got %d entries, want 6", len(got))
	}
	if got[0].Level != "ERROR" || got[0].Message != "disk full" {
		t.Errorf("got[0] = %+v, want {ERROR disk full}", got[0])
	}
}

func TestFilterLevel(t *testing.T) {
	entries := Parse(strings.NewReader(testLog))
	errors := FilterLevel(entries, "ERROR")
	count := 0
	for range errors {
		count++
	}
	if count != 2 {
		t.Errorf("FilterLevel(ERROR) count = %d, want 2", count)
	}
}

func TestCountByLevel(t *testing.T) {
	entries := Parse(strings.NewReader(testLog))
	counts := CountByLevel(entries)
	if counts["ERROR"] != 2 {
		t.Errorf("ERROR count = %d, want 2", counts["ERROR"])
	}
	if counts["INFO"] != 2 {
		t.Errorf("INFO count = %d, want 2", counts["INFO"])
	}
}
