package exercise_05

import "fmt"

// Notifier sends notifications.
type Notifier interface {
	Notify(message string) error
}

// EmailNotifier sends emails (simulated).
type EmailNotifier struct {
	Address string
}

// TODO: implement Notify — return nil, set Sent field to true, record message
func (e *EmailNotifier) Notify(message string) error {
	return fmt.Errorf("not implemented")
}

// SlackNotifier sends Slack messages (simulated).
type SlackNotifier struct {
	Channel string
}

// TODO: implement Notify — return nil, record message
func (s *SlackNotifier) Notify(message string) error {
	return fmt.Errorf("not implemented")
}

// BroadcastAll sends message to all notifiers.
// Returns the first error encountered, or nil.
// TODO: iterate and call n.Notify(message)
func BroadcastAll(notifiers []Notifier, message string) error {
	return fmt.Errorf("not implemented")
}

// FailingNotifier always returns an error — used for testing.
type FailingNotifier struct{}

func (f FailingNotifier) Notify(_ string) error {
	return fmt.Errorf("notification failed")
}
