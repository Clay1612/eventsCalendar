package calendar

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/Clay1612/eventsCalendar/events"
	"github.com/Clay1612/eventsCalendar/helpers"
	"github.com/Clay1612/eventsCalendar/storage"
)

type Calendar struct {
	CalendarEvents map[string]*events.Event
	Storage        storage.Store
	Notification   chan string
}

func NewCalendar(store storage.Store) *Calendar {
	return &Calendar{
		CalendarEvents: make(map[string]*events.Event),
		Storage:        store,
		Notification:   make(chan string),
	}
}

func (c *Calendar) GetEvents() map[string]*events.Event {
	return c.CalendarEvents
}

func (c *Calendar) AddEvent(title string, date string, priority events.Priority) (*events.Event, error) {
	for _, event := range c.CalendarEvents {
		if event.Title == title {
			return event, nil
		}
	}

	e, err := events.NewEvent(title, date, priority)

	if err != nil {
		return &events.Event{}, fmt.Errorf("add event func error: %w", err)
	}

	c.CalendarEvents[e.ID] = e

	return e, nil
}

func (c *Calendar) DeleteEvent(name string) error {
	for _, event := range c.CalendarEvents {
		if event.Title == name {
			delete(c.CalendarEvents, event.ID)
			return nil
		}
	}

	return fmt.Errorf("delete event func error: %w", helpers.KnownErrors["EventNotFoundError"])
}

func (c *Calendar) EditEvent(name string, dateStr string, priority events.Priority) error {
	for _, event := range c.CalendarEvents {
		if event.Title == name {
			err := event.UpdateEvent(name, dateStr, priority)
			if err != nil {
				return fmt.Errorf("edit event func error: %w", err)
			}

			return nil
		}
	}

	return fmt.Errorf("edit event func error: %w", helpers.KnownErrors["EventNotFoundError"])
}

func (c *Calendar) SetEventReminder(eventName string, dateStr string, reminderMsg string) error {
	for _, event := range c.CalendarEvents {
		if event.Title == eventName {
			delayMinutes, err := strconv.Atoi(dateStr)
			if err != nil {
				return fmt.Errorf("set event reminder func error: %w", helpers.KnownErrors["StrconvAtoiError"])
			}

			reminderTime := event.StartAt.Add(time.Duration(-delayMinutes) * time.Minute)
			timerDuration := reminderTime.Sub(time.Now())

			err = event.AddReminder(reminderMsg, timerDuration, c.Notify)
			if err != nil {
				return fmt.Errorf("set event reminder func error: %w", err)
			}

			return nil
		}
	}

	return fmt.Errorf("set event reminder func error: %w", helpers.KnownErrors["EventNotFoundError"])
}

func (c *Calendar) CancelEventReminder(eventName string) error {
	for _, event := range c.CalendarEvents {
		if event.Title == eventName {
			if event.Reminder == nil {
				return fmt.Errorf("cancel event reminder error: %w", helpers.KnownErrors["ReminderAlreadyRemovedError"])
			}

			event.RemoveReminder()

			return nil
		}
	}

	return fmt.Errorf("cancel event reminder error:: %w", helpers.KnownErrors["EventNotFoundError"])
}

func (c *Calendar) Notify(msg string) {
	c.Notification <- msg
}

func (c *Calendar) Save() error {
	data, err := json.Marshal(c.CalendarEvents)

	if err != nil {
		return fmt.Errorf("save func error: %w", err)
	}

	err = c.Storage.Save(data)
	if err != nil {
		return fmt.Errorf("save func error: %w", err)
	}

	return nil
}

func (c *Calendar) Load() error {
	data, err := c.Storage.Load()

	if err != nil {
		return fmt.Errorf("load func error: %w", err)
	}

	err = json.Unmarshal(data, &c.CalendarEvents)
	if err != nil {
		return fmt.Errorf("load func error: %w", err)
	}

	return nil
}
