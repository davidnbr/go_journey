package exercise_03

import "fmt"

// Animal has a name and sound.
type Animal struct {
	Name string
}

// Speak returns "<Name> says ..."
// TODO: implement
func (a Animal) Speak() string {
	return ""
}

// Dog embeds Animal and adds a breed.
type Dog struct {
	Animal
	Breed string
}

// Fetch returns "<Name> fetches the ball!"
// TODO: access the embedded Name field directly
func (d Dog) Fetch() string {
	return ""
}

// Logger embeds a prefix string for log messages.
type Logger struct {
	Prefix string
}

// Log formats a message: "[Prefix] message"
// TODO: use fmt.Sprintf
func (l Logger) Log(message string) string {
	return ""
}

// Service embeds Logger and has a name.
type Service struct {
	Logger
	Name string
}

var _ = fmt.Sprintf
