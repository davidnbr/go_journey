package exercise_07

import (
	"errors"
	"testing"
)

func TestValidationError(t *testing.T) {
	e := &ValidationError{Field: "email", Message: "invalid format"}
	s := e.Error()
	if s == "" {
		t.Error("ValidationError.Error() must not be empty")
	}
}

func TestValidateAge(t *testing.T) {
	if err := ValidateAge(25); err != nil {
		t.Errorf("ValidateAge(25) unexpected error: %v", err)
	}
	err := ValidateAge(-1)
	if err == nil {
		t.Fatal("ValidateAge(-1) should return error")
	}
	var ve *ValidationError
	if !errors.As(err, &ve) {
		t.Errorf("ValidateAge(-1) should return *ValidationError, got %T", err)
	}
	if ve.Field != "age" {
		t.Errorf("ValidationError.Field = %q, want \"age\"", ve.Field)
	}
	err = ValidateAge(200)
	if err == nil {
		t.Error("ValidateAge(200) should return error")
	}
}

func TestFindUser(t *testing.T) {
	users := map[string]string{"u1": "Alice", "u2": "Bob"}
	name, err := FindUser(users, "u1")
	if err != nil || name != "Alice" {
		t.Errorf("FindUser(u1) = (%q, %v), want (Alice, nil)", name, err)
	}
	_, err = FindUser(users, "u99")
	if err == nil {
		t.Fatal("FindUser(u99) should return error")
	}
	var nfe *NotFoundError
	if !errors.As(err, &nfe) {
		t.Errorf("FindUser(u99) should return *NotFoundError, got %T", err)
	}
}
