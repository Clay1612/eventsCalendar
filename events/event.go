package events

import (
	"fmt"
	"time"

	"github.com/Clay1612/eventsCalendar/helpers"
	"github.com/Clay1612/eventsCalendar/reminder"
	"github.com/Clay1612/eventsCalendar/validation"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID       string             `json:"id"`
	Title    string             `json:"title"`
	StartAt  time.Time          `json:"start_at"`
	Priority Priority           `json:"priority"`
	Reminder *reminder.Reminder `json:"reminder"`
}

type Priority string

func (p Priority) Validate() error {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	default:
		return helpers.KnownErrors["PriorityValidationError"]
	}
}

const (
	PriorityLow    Priority = "low"
	PriorityMedium Priority = "medium"
	PriorityHigh   Priority = "high"
)

func GetNewID() string {
	return uuid.New().String()
}

func NewEvent(title string, dateStr string, priority Priority) (*Event, error) {
	var eventTime time.Time
	var isValidTitle bool
	var err error

	isValidTitle, err = validation.IsValidString(title)

	if err != nil {
		return &Event{}, fmt.Errorf("new event func error: %w", err)
	}

	if !isValidTitle {
		return &Event{}, fmt.Errorf("new event func error: %w", helpers.KnownErrors["TitleValidationError"])
	}

	if err = priority.Validate(); err != nil {
		return &Event{}, fmt.Errorf("new event func error: %w", err)
	}

	eventTime, err = dateparse.ParseLocal(dateStr)

	if err != nil {
		return &Event{}, fmt.Errorf("new event func error: %w", helpers.KnownErrors["ParseDateError"])
	}

	return &Event{
		ID:       GetNewID(),
		Title:    title,
		StartAt:  eventTime,
		Priority: priority,
		Reminder: nil,
	}, nil
}

func (e *Event) UpdateEvent(title string, dateStr string, priority Priority) error {
	var newTime time.Time
	var isValidTitle bool
	var err error

	isValidTitle, err = validation.IsValidString(title)

	if err != nil {
		return fmt.Errorf("update func error: %w", err)
	}

	if !isValidTitle {
		return fmt.Errorf("update func error: %w", helpers.KnownErrors["TitleValidationError"])
	}

	if err = priority.Validate(); err != nil {
		return fmt.Errorf("update func error: %w", err)
	}

	newTime, err = dateparse.ParseLocal(dateStr)

	if err != nil {
		return fmt.Errorf("update func error: %w", helpers.KnownErrors["ParseDateError"])
	}

	e.Title = title
	e.StartAt = newTime
	e.Priority = priority

	return nil
}

func (e *Event) AddReminder(msg string, reminderDuration time.Duration, notifier func(msg string)) error {
	if e.Reminder != nil {
		return fmt.Errorf("add reminder func error: %w", helpers.KnownErrors["ReminderAddingError"])
	}

	r, err := reminder.NewReminder(msg, reminderDuration, notifier)

	if err != nil {
		return fmt.Errorf("add reminder func error: %w", err)
	}

	e.Reminder = r
	e.Reminder.Start()

	return nil
}

func (e *Event) RemoveReminder() {
	if e.Reminder.Timer != nil {
		e.Reminder.Stop()
	}

	e.Reminder = nil
}
