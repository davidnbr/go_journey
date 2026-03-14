package exercise_05

import "testing"

func TestEmailNotifier(t *testing.T) {
	e := &EmailNotifier{Address: "ops@example.com"}
	if err := e.Notify("deploy started"); err != nil {
		t.Errorf("EmailNotifier.Notify() error: %v", err)
	}
}

func TestSlackNotifier(t *testing.T) {
	s := &SlackNotifier{Channel: "#alerts"}
	if err := s.Notify("CPU high"); err != nil {
		t.Errorf("SlackNotifier.Notify() error: %v", err)
	}
}

func TestBroadcastAll(t *testing.T) {
	notifiers := []Notifier{
		&EmailNotifier{Address: "a@example.com"},
		&SlackNotifier{Channel: "#ops"},
	}
	if err := BroadcastAll(notifiers, "hello"); err != nil {
		t.Errorf("BroadcastAll unexpected error: %v", err)
	}
}

func TestBroadcastAllError(t *testing.T) {
	notifiers := []Notifier{
		&EmailNotifier{Address: "a@example.com"},
		FailingNotifier{},
	}
	if err := BroadcastAll(notifiers, "hello"); err == nil {
		t.Error("BroadcastAll should return error when a notifier fails")
	}
}

func TestInterfaceSatisfaction(t *testing.T) {
	var _ Notifier = &EmailNotifier{}
	var _ Notifier = &SlackNotifier{}
	var _ Notifier = FailingNotifier{}
}
