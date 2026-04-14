package exercise_14

import (
	"fmt"
	"reflect"
)

// TypeName returns the type name of v as a string.
// e.g., int, string, []int, map[string]int
// TODO: use reflect.TypeOf(v).String()
func TypeName(v interface{}) string {
	return ""
}

// FieldNames returns the exported field names of a struct.
// Returns nil for non-struct types.
// TODO: reflect.TypeOf, check Kind, iterate NumField
func FieldNames(v interface{}) []string {
	return nil
}

// SetField sets the field named fieldName on ptr (a pointer to a struct) to value.
// Returns error if field not found or not settable.
// TODO: reflect.ValueOf(ptr).Elem(), FieldByName, Set
func SetField(ptr interface{}, fieldName string, value interface{}) error {
	return fmt.Errorf("not implemented")
}

var _ = reflect.TypeOf
