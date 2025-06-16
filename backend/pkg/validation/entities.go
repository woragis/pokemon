package validation

import (
	"errors"

	"github.com/google/uuid"
)

func UUIDRequired(id uuid.UUID, field string) error {
	if id == uuid.Nil {
		return errors.New(field + " is required")
	}
	return nil
}

func StringRequired(str string, field string) error {
	if str == "" {
		return errors.New(field + " is required")
	}
	return nil
}

func StringMaxLength(str string, max int, field string) error {
	if len(str) > max {
		return errors.New(field + " must be less than " + string(rune(max)) + " characters")
	}
	return nil
}
