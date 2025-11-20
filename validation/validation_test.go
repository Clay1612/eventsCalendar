package validation

import (
	"testing"
)

func TestIsValidString(t *testing.T) {
	trueString := "Hello, World!"
	trueValid, err := IsValidString(trueString)
	if err != nil {
		t.Error(err)
	}

	if !trueValid {
		t.Errorf("validation - %v. Expected valid - true", trueValid)
	}

	falseString := "Hello?"
	falseValid, err2 := IsValidString(falseString)
	if err2 != nil {
		t.Error(err2)
	}
	if falseValid {
		t.Errorf("validation - %v. Expected invalid - false", falseValid)
	}
}
