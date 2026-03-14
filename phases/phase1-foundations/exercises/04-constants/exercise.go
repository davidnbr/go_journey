package exercise_04

// Direction represents a compass direction using iota.
// TODO: define Direction as int type, then declare North=iota, East, South, West
type Direction int

const (
	North Direction = iota
	East
	South
	West
)

// Pi is the mathematical constant π to 5 decimal places.
// TODO: define Pi as an untyped constant = 3.14159
const Pi = 0.0

// DaysInWeek is the number of days in a week.
// TODO: define as typed int constant = 7
const DaysInWeek int = 0

// DirectionName returns the string name of a Direction.
// TODO: use a switch to return "North", "East", "South", "West", or "Unknown"
func DirectionName(d Direction) string {
	return ""
}
