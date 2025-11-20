package validation

import (
	"fmt"
	"regexp"

	"github.com/Clay1612/eventsCalendar/helpers"
)

func IsValidString(str string) (bool, error) {
	pattern := "^[a-zA-Z0-9 !,/.']{3,50}$"
	matched, err := regexp.MatchString(pattern, str)

	if err != nil {
		return matched, fmt.Errorf("regexp.MatchString func err: %w", helpers.KnownErrors["MatchStringError"])
	}

	return matched, nil
}
