package exercise_04

import "testing"

func TestIota(t *testing.T) {
	if North != 0 {
		t.Errorf("North = %d, want 0", North)
	}
	if East != 1 {
		t.Errorf("East = %d, want 1", East)
	}
	if South != 2 {
		t.Errorf("South = %d, want 2", South)
	}
	if West != 3 {
		t.Errorf("West = %d, want 3", West)
	}
}

func TestPi(t *testing.T) {
	if Pi != 3.14159 {
		t.Errorf("Pi = %v, want 3.14159", Pi)
	}
}

func TestDaysInWeek(t *testing.T) {
	if DaysInWeek != 7 {
		t.Errorf("DaysInWeek = %d, want 7", DaysInWeek)
	}
}

func TestDirectionName(t *testing.T) {
	tests := []struct {
		d    Direction
		want string
	}{
		{North, "North"},
		{East, "East"},
		{South, "South"},
		{West, "West"},
		{Direction(99), "Unknown"},
	}
	for _, tt := range tests {
		got := DirectionName(tt.d)
		if got != tt.want {
			t.Errorf("DirectionName(%d) = %q, want %q", tt.d, got, tt.want)
		}
	}
}
