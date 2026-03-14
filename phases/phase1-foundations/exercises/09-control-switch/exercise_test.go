package exercise_09

import "testing"

func TestDayType(t *testing.T) {
	tests := []struct{ day, want string }{
		{"Monday", "Weekday"},
		{"Tuesday", "Weekday"},
		{"Wednesday", "Weekday"},
		{"Thursday", "Weekday"},
		{"Friday", "Weekday"},
		{"Saturday", "Weekend"},
		{"Sunday", "Weekend"},
		{"Holiday", "Unknown"},
	}
	for _, tt := range tests {
		got := DayType(tt.day)
		if got != tt.want {
			t.Errorf("DayType(%q) = %q, want %q", tt.day, got, tt.want)
		}
	}
}

func TestHTTPStatus(t *testing.T) {
	tests := []struct {
		code int
		want string
	}{
		{200, "OK"},
		{201, "Created"},
		{301, "Moved Permanently"},
		{400, "Bad Request"},
		{401, "Unauthorized"},
		{403, "Forbidden"},
		{404, "Not Found"},
		{500, "Internal Server Error"},
		{418, "Unknown Status"},
	}
	for _, tt := range tests {
		got := HTTPStatus(tt.code)
		if got != tt.want {
			t.Errorf("HTTPStatus(%d) = %q, want %q", tt.code, got, tt.want)
		}
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		v    interface{}
		want string
	}{
		{42, "int"},
		{"hello", "string"},
		{true, "bool"},
		{3.14, "float64"},
		{[]int{}, "unknown"},
	}
	for _, tt := range tests {
		got := Describe(tt.v)
		if got != tt.want {
			t.Errorf("Describe(%v) = %q, want %q", tt.v, got, tt.want)
		}
	}
}
