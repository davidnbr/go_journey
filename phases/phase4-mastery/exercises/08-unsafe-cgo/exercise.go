package exercise_08

import (
	"unsafe"
)

// SizeOf returns the size in bytes of a value using unsafe.Sizeof.
// TODO: return int(unsafe.Sizeof(v))
func SizeOf(v interface{}) int {
	_ = unsafe.Sizeof(v) // suppress unused import
	return 0 // stub
}

// AlignOf returns the alignment of a value.
// TODO: return int(unsafe.Alignof(v))
func AlignOf(v interface{}) int {
	return 0 // stub
}

// StringToBytes converts a string to []byte without allocation.
// WARNING: The returned slice must not be modified. Educational only.
// TODO: use unsafe.StringData / unsafe pointers
func StringToBytes(s string) []byte {
	return []byte(s) // stub: safe copy (not zero-alloc)
}
