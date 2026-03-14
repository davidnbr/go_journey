package wordcounter

import (
	"strings"
	"testing"
)

func TestCountWords(t *testing.T) {
	input := "go is great go is fast go"
	freq, err := CountWords(strings.NewReader(input))
	if err != nil {
		t.Fatalf("CountWords error: %v", err)
	}
	if freq["go"] != 3 {
		t.Errorf("go count = %d, want 3", freq["go"])
	}
	if freq["is"] != 2 {
		t.Errorf("is count = %d, want 2", freq["is"])
	}
	if freq["great"] != 1 {
		t.Errorf("great count = %d, want 1", freq["great"])
	}
	if freq["fast"] != 1 {
		t.Errorf("fast count = %d, want 1", freq["fast"])
	}
}

func TestCountWordsCaseInsensitive(t *testing.T) {
	freq, err := CountWords(strings.NewReader("Go GO go"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if freq["go"] != 3 {
		t.Errorf("case-insensitive count: go=%d, want 3", freq["go"])
	}
}

func TestTopN(t *testing.T) {
	freq := map[string]int{"go": 5, "is": 3, "fun": 3, "fast": 1}
	top := TopN(freq, 2)
	if len(top) != 2 {
		t.Fatalf("TopN(2) len = %d, want 2", len(top))
	}
	if top[0].Word != "go" || top[0].Count != 5 {
		t.Errorf("top[0] = %+v, want {go 5}", top[0])
	}
	// "fun" and "is" both have count 3 — sorted alphabetically: "fun" < "is"
	if top[1].Word != "fun" {
		t.Errorf("top[1].Word = %q, want \"fun\" (alpha tie-break)", top[1].Word)
	}
}

func TestTopNAll(t *testing.T) {
	freq := map[string]int{"a": 1, "b": 2}
	all := TopN(freq, 0)
	if len(all) != 2 {
		t.Errorf("TopN(0) should return all %d entries, got %d", len(freq), len(all))
	}
}
