package exercise_07

import "fmt"

// ValidationError is a custom error type with a field name and message.
type ValidationError struct {
	Field   string
	Message string
}

// TODO: implement the error interface — return "field <Field>: <Message>"
func (e *ValidationError) Error() string {
	return ""
}

// NotFoundError represents a missing resource.
type NotFoundError struct {
	Resource string
	ID       string
}

// TODO: implement error interface — return "<Resource> with id <ID> not found"
func (e *NotFoundError) Error() string {
	return ""
}

// ValidateAge returns a *ValidationError if age is outside [0, 150].
// Returns nil for valid age.
// TODO: check bounds, return &ValidationError{Field: "age", Message: "..."}
func ValidateAge(age int) error {
	return nil
}

// FindUser returns a *NotFoundError if the id is not in users.
// Returns nil error if found.
// TODO: check map, return &NotFoundError{Resource: "user", ID: id}
func FindUser(users map[string]string, id string) (string, error) {
	return "", nil
}

var _ = fmt.Sprintf
