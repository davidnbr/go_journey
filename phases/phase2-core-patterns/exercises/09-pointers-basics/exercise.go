package exercise_09

// Increment adds 1 to the value at the pointer.
// TODO: dereference and increment: *n++
func Increment(n *int) {
	// no-op stub
}

// Swap exchanges the values of the two pointers.
// TODO: dereference swap: *a, *b = *b, *a
func Swap(a, b *int) {
	// no-op stub
}

// NewInt allocates a new int on the heap with the given value and returns its address.
// TODO: use new(int) or &value pattern
func NewInt(v int) *int {
	return nil
}

// SafeDeref returns the value at ptr, or defaultVal if ptr is nil.
// TODO: nil check, then dereference
func SafeDeref(ptr *int, defaultVal int) int {
	return 0
}
