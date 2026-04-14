package exercise_04

import "fmt"

// EmailSender sends emails.
type EmailSender interface {
	Send(to, subject, body string) error
}

// UserService manages users and sends welcome emails.
type UserService struct {
	sender EmailSender
}

// NewUserService creates a UserService with the given sender.
func NewUserService(sender EmailSender) *UserService {
	return &UserService{sender: sender}
}

// RegisterUser creates a user and sends a welcome email.
// Returns an error if the email or sending fails.
// Returns error if name or email is empty.
// TODO: validate inputs, send welcome email via s.sender
func (s *UserService) RegisterUser(name, email string) error {
	return fmt.Errorf("not implemented")
}
