package exercise_03

import (
	"strings"
	"testing"
)

func TestAnimalSpeak(t *testing.T) {
	a := Animal{Name: "Cat"}
	s := a.Speak()
	if !strings.Contains(s, "Cat") {
		t.Errorf("Speak() = %q, want string containing Cat", s)
	}
}

func TestDogEmbedding(t *testing.T) {
	d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
	// Promoted field access
	if d.Name != "Rex" {
		t.Errorf("d.Name = %q, want Rex", d.Name)
	}
	// Promoted method
	if s := d.Speak(); !strings.Contains(s, "Rex") {
		t.Errorf("d.Speak() = %q, want string containing Rex", s)
	}
	// Dog-specific method
	fetch := d.Fetch()
	if !strings.Contains(fetch, "Rex") {
		t.Errorf("Fetch() = %q, want string containing Rex", fetch)
	}
}

func TestServiceLogger(t *testing.T) {
	svc := Service{Logger: Logger{Prefix: "api"}, Name: "AuthService"}
	// Promoted Log method
	msg := svc.Log("started")
	if !strings.Contains(msg, "api") || !strings.Contains(msg, "started") {
		t.Errorf("svc.Log() = %q, want [api] started", msg)
	}
}
