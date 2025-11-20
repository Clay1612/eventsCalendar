package reminder

import (
	"fmt"
	"time"

	"github.com/Clay1612/eventsCalendar/helpers"
	"github.com/Clay1612/eventsCalendar/validation"
)

type Reminder struct {
	Message string
	At      time.Duration
	Sent    bool
	Timer   *time.Timer      `json:"-"`
	Notify  func(msg string) `json:"-"`
}

func NewReminder(msg string, t time.Duration, notifier func(msg string)) (*Reminder, error) {
	isValid, err := validation.IsValidString(msg)

	if err != nil {
		return &Reminder{}, err
	}

	if !isValid {
		return &Reminder{}, fmt.Errorf("new reminder func error: %w", helpers.KnownErrors["InvalidMessageErr"])
	}

	return &Reminder{
		Message: msg,
		At:      t,
		Sent:    false,
		Timer:   nil,
		Notify:  notifier,
	}, nil
}

func (r *Reminder) Send() {
	if r.Sent {
		return
	}

	r.Notify(r.Message)
	r.Sent = true
}

func (r *Reminder) Start() {
	r.Timer = time.AfterFunc(r.At, r.Send)
}

func (r *Reminder) Stop() {
	r.Timer.Stop()
}
