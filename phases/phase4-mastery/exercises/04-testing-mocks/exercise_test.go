package exercise_04

import (
	"errors"
	"testing"
)

// MockEmailSender is a test double that records calls.
type MockEmailSender struct {
	Calls []EmailCall
	ErrOn string // if To matches this, return an error
}

type EmailCall struct {
	To, Subject, Body string
}

func (m *MockEmailSender) Send(to, subject, body string) error {
	m.Calls = append(m.Calls, EmailCall{To: to, Subject: subject, Body: body})
	if m.ErrOn == to {
		return errors.New("send failed")
	}
	return nil
}

func TestRegisterUser(t *testing.T) {
	mock := &MockEmailSender{}
	svc := NewUserService(mock)

	err := svc.RegisterUser("Alice", "alice@example.com")
	if err != nil {
		t.Fatalf("RegisterUser error: %v", err)
	}
	if len(mock.Calls) != 1 {
		t.Fatalf("expected 1 email, got %d", len(mock.Calls))
	}
	if mock.Calls[0].To != "alice@example.com" {
		t.Errorf("email To = %q, want alice@example.com", mock.Calls[0].To)
	}
}

func TestRegisterUserValidation(t *testing.T) {
	mock := &MockEmailSender{}
	svc := NewUserService(mock)

	if err := svc.RegisterUser("", "a@b.com"); err == nil {
		t.Error("empty name should return error")
	}
	if err := svc.RegisterUser("Alice", ""); err == nil {
		t.Error("empty email should return error")
	}
	if len(mock.Calls) != 0 {
		t.Error("no email should be sent on validation failure")
	}
}

func TestRegisterUserSendFailure(t *testing.T) {
	mock := &MockEmailSender{ErrOn: "fail@example.com"}
	svc := NewUserService(mock)

	err := svc.RegisterUser("Bob", "fail@example.com")
	if err == nil {
		t.Error("should return error when send fails")
	}
}
