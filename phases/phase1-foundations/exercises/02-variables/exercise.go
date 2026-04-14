package exercise_02

// ZeroValues returns the zero values of int, float64, bool, and string.
// TODO: return 0, 0.0, false, ""  (they are already zero values — just declare and return)
func ZeroValues() (int, float64, bool, string) {
	return 1, 1.0, true, "non-empty"
}

// Swap returns two strings with their values exchanged.
// TODO: return b, a
func Swap(a, b string) (string, string) {
	return a, b
}

// ShortDeclaration uses := to declare and return a sum.
// TODO: use := to assign x = 10, y = 20, return x + y
func ShortDeclaration() int {
	return 0
}
