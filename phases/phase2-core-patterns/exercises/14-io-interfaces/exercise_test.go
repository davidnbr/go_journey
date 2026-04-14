package exercise_14

import (
	"bytes"
	"strings"
	"testing"
)

func TestReadAll(t *testing.T) {
	r := strings.NewReader("hello world")
	got, err := ReadAll(r)
	if err != nil {
		t.Fatalf("ReadAll error: %v", err)
	}
	if got != "hello world" {
		t.Errorf("ReadAll = %q, want \"hello world\"", got)
	}
}

func TestWriteString(t *testing.T) {
	var buf bytes.Buffer
	err := WriteString(&buf, "hello")
	if err != nil {
		t.Fatalf("WriteString error: %v", err)
	}
	if buf.String() != "hello" {
		t.Errorf("WriteString wrote %q, want \"hello\"", buf.String())
	}
}

func TestUpperReader(t *testing.T) {
	r := strings.NewReader("hello world")
	upper, err := UpperReader(r)
	if err != nil {
		t.Fatalf("UpperReader error: %v", err)
	}
	if upper == nil {
		t.Fatal("UpperReader returned nil reader")
	}
	result, err := ReadAll(upper)
	if err != nil {
		t.Fatalf("ReadAll after UpperReader error: %v", err)
	}
	if result != "HELLO WORLD" {
		t.Errorf("UpperReader result = %q, want \"HELLO WORLD\"", result)
	}
}
