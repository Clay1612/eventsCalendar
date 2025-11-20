package events

import (
	"testing"
)

func TestValidate(t *testing.T) {
	var truePriority Priority = "medium"

	err := truePriority.Validate()
	if err != nil {
		t.Error(err)
	}

	var falsePriority Priority = "urgent"
	err2 := falsePriority.Validate()
	if err2 == nil {
		t.Error("Expected an error")
	}
}

func TestNewEvent(t *testing.T) {
	var (
		trueTitle              = "hello world"
		trueDateStr            = "2025-01-01"
		truePriority  Priority = "medium"
		falseTitle             = "Здраствуй мир?"
		falseDateStr           = "01-13-2025"
		falsePriority Priority = "urgent"
	)

	_, err := NewEvent(trueTitle, trueDateStr, truePriority)
	if err != nil {
		t.Error(err)
	}

	_, err = NewEvent(falseTitle, trueDateStr, truePriority)
	if err == nil {
		t.Error("Expected an error in title")
	}

	_, err = NewEvent(trueTitle, falseDateStr, truePriority)
	if err == nil {
		t.Error("Expected an error in date")
	}

	_, err = NewEvent(trueTitle, trueDateStr, falsePriority)
	if err == nil {
		t.Error("Expected an error in priority")
	}
}

func TestUpdateEvent(t *testing.T) {
	var (
		trueTitle                   = "hello world"
		trueDateStr                 = "2025-01-01"
		truePriority       Priority = "medium"
		falseTitle                  = "Здраствуй мир?"
		falseDateStr                = "01-13-2025"
		falsePriority      Priority = "urgent"
		secondTrueTitle             = "write some code"
		secondTrueDateStr           = "2026-01-01"
		secondTruePriority Priority = "high"
	)

	event, _ := NewEvent(trueTitle, trueDateStr, truePriority)

	err := event.UpdateEvent(secondTrueTitle, secondTrueDateStr, secondTruePriority)
	if err != nil {
		t.Error(err)
	}

	err = event.UpdateEvent(falseTitle, secondTrueDateStr, secondTruePriority)
	if err == nil {
		t.Error("Expected an error in title")
	}

	err = event.UpdateEvent(secondTrueTitle, falseDateStr, secondTruePriority)
	if err == nil {
		t.Error("Expected an error in date")
	}

	err = event.UpdateEvent(secondTrueTitle, secondTrueDateStr, falsePriority)
	if err == nil {
		t.Error("Expected an error in priority")
	}
}
