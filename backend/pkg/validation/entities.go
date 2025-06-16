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

func Email() error {
	return nil
}

func Password() error {
	return nil
}

func Username() error {
	return nil
}

func DateOfBirth() error {
	return nil
}

func CreditCard() error {
	return nil
}

func Phone() error {
	return nil
}

func Timezone() error {
	return nil
}

func FileExtension() error {
	return nil
}

func Urls() error {
	return nil
}

func Address() error {
	return nil
}
