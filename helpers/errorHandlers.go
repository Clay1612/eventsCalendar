package helpers

import (
	"errors"
	"fmt"
)

var KnownErrors = map[string]error{
	"MatchStringError":            errors.New("invalid regular expression or pattern"),
	"InvalidMessageErr":           errors.New("invalid message"),
	"TitleValidationError":        errors.New("the title is not valid"),
	"PriorityValidationError":     errors.New("invalid Priority"),
	"ParseDateError":              errors.New("failed to parse the date"),
	"ReminderAddingError":         errors.New("the reminder has already been added"),
	"EventNotFoundError":          errors.New("event not found"),
	"ReminderAlreadyRemovedError": errors.New("reminder already removed"),
	"StrconvAtoiError":            errors.New("strconv atoi func error"),
}

func ErrorHandler(wrappedError error) {
	for _, e := range KnownErrors {
		if errors.Is(wrappedError, e) {
			fmt.Println(e)
			return
		}
	}

	fmt.Println(wrappedError)
}
