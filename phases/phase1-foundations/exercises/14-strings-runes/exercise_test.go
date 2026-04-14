package exercise_14

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"racecar", true},
		{"Racecar", true},
		{"hello", false},
		{"", true},
		{"a", true},
		{"Aba", true},
	}
	for _, tt := range tests {
		got := IsPalindrome(tt.s)
		if got != tt.want {
			t.Errorf("IsPalindrome(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}

func TestCountRunes(t *testing.T) {
	// "Hello" = 5 bytes = 5 runes
	if got := CountRunes("Hello"); got != 5 {
		t.Errorf("CountRunes(Hello) = %d, want 5", got)
	}
	// "日本語" = 9 bytes, 3 runes
	if got := CountRunes("日本語"); got != 3 {
		t.Errorf("CountRunes(日本語) = %d, want 3", got)
	}
}

func TestReverseString(t *testing.T) {
	if got := ReverseString("hello"); got != "olleh" {
		t.Errorf("ReverseString(hello) = %q, want \"olleh\"", got)
	}
	if got := ReverseString("日本語"); got != "語本日" {
		t.Errorf("ReverseString(日本語) = %q, want \"語本日\"", got)
	}
}

func TestContainsAll(t *testing.T) {
	if !ContainsAll("hello world", []string{"hello", "world"}) {
		t.Error("ContainsAll should return true")
	}
	if ContainsAll("hello world", []string{"hello", "foo"}) {
		t.Error("ContainsAll should return false when one sub is missing")
	}
	if !ContainsAll("anything", []string{}) {
		t.Error("ContainsAll with empty subs should return true")
	}
}
