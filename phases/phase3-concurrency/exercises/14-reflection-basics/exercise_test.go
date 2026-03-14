package exercise_14

import (
	"testing"
)

type Person struct {
	Name string
	Age  int
}

func TestTypeName(t *testing.T) {
	tests := []struct {
		v    interface{}
		want string
	}{
		{42, "int"},
		{"hello", "string"},
		{[]int{}, "[]int"},
		{map[string]int{}, "map[string]int"},
	}
	for _, tt := range tests {
		got := TypeName(tt.v)
		if got != tt.want {
			t.Errorf("TypeName(%v) = %q, want %q", tt.v, got, tt.want)
		}
	}
}

func TestFieldNames(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	names := FieldNames(p)
	if len(names) != 2 {
		t.Fatalf("FieldNames(Person) len = %d, want 2", len(names))
	}
	if names[0] != "Name" || names[1] != "Age" {
		t.Errorf("FieldNames = %v, want [Name Age]", names)
	}
	if FieldNames(42) != nil {
		t.Error("FieldNames(non-struct) should return nil")
	}
}

func TestSetField(t *testing.T) {
	p := &Person{}
	if err := SetField(p, "Name", "Bob"); err != nil {
		t.Fatalf("SetField Name: %v", err)
	}
	if p.Name != "Bob" {
		t.Errorf("p.Name = %q, want Bob", p.Name)
	}
	if err := SetField(p, "Age", 25); err != nil {
		t.Fatalf("SetField Age: %v", err)
	}
	if p.Age != 25 {
		t.Errorf("p.Age = %d, want 25", p.Age)
	}
	if err := SetField(p, "Missing", "x"); err == nil {
		t.Error("SetField(missing) should return error")
	}
}
