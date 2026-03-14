package exercise_13

import "testing"

func TestModuleInfo(t *testing.T) {
	got := ModuleInfo()
	if got != "github.com/davidnbr/go_journey" {
		t.Errorf("ModuleInfo() = %q, want github.com/davidnbr/go_journey", got)
	}
}

func TestGoVersion(t *testing.T) {
	got := GoVersion()
	if got != "1.22" {
		t.Errorf("GoVersion() = %q, want 1.22", got)
	}
}

func TestIsSemanticVersion(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"v1.2.3", true},
		{"v0.0.1", true},
		{"v1.22.0", true},
		{"1.2.3", false},
		{"v1.2", false},
		{"latest", false},
	}
	for _, tt := range tests {
		got := IsSemanticVersion(tt.s)
		if got != tt.want {
			t.Errorf("IsSemanticVersion(%q) = %v, want %v", tt.s, got, tt.want)
		}
	}
}
