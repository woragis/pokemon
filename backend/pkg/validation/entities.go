package validation

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"

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

// Validate email format
func Email(value string) error {
	_, err := mail.ParseAddress(value)
	if err != nil {
		return errors.New("invalid email format")
	}
	return nil
}

// Validate password: min 8 chars, at least one uppercase, one number, one special char
func Password(value string) error {
	if len(value) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(value) {
		return errors.New("password must contain at least one uppercase letter")
	}
	if !regexp.MustCompile(`[0-9]`).MatchString(value) {
		return errors.New("password must contain at least one number")
	}
	if !regexp.MustCompile(`[!@#\$%\^&\*]`).MatchString(value) {
		return errors.New("password must contain at least one special character")
	}
	return nil
}

// Validate username: 3-20 chars, letters, numbers, underscores only
func Username(value string) error {
	if len(value) < 3 || len(value) > 20 {
		return errors.New("username must be between 3 and 20 characters")
	}
	if !regexp.MustCompile(`^[a-zA-Z0-9_]+$`).MatchString(value) {
		return errors.New("username can only contain letters, numbers, and underscores")
	}
	return nil
}

// Validate date of birth: must be a valid date in the past, and at least 13 years old
func DateOfBirth(value string) error {
	dob, err := time.Parse("2006-01-02", value) // expect format YYYY-MM-DD
	if err != nil {
		return errors.New("invalid date format, expected YYYY-MM-DD")
	}
	if dob.After(time.Now()) {
		return errors.New("date of birth cannot be in the future")
	}
	age := time.Now().Year() - dob.Year()
	if age < 13 {
		return errors.New("must be at least 13 years old")
	}
	return nil
}

// Validate credit card number (basic Luhn check)
func CreditCard(value string) error {
	value = strings.ReplaceAll(value, " ", "") // remove spaces
	if len(value) < 13 || len(value) > 19 {
		return errors.New("credit card number length invalid")
	}
	if !luhnCheck(value) {
		return errors.New("invalid credit card number")
	}
	return nil
}

// Luhn algorithm implementation
func luhnCheck(number string) bool {
	sum := 0
	double := false
	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}

// Validate phone number (simple regex for digits, may customize)
func Phone(value string) error {
	// Allow digits, spaces, +, -, ()
	if !regexp.MustCompile(`^[\d\s\+\-\(\)]+$`).MatchString(value) {
		return errors.New("invalid phone number format")
	}
	return nil
}

// Validate timezone string (checks if valid IANA tz)
func Timezone(value string) error {
	_, err := time.LoadLocation(value)
	if err != nil {
		return errors.New("invalid timezone")
	}
	return nil
}

// Validate file extension (e.g. ".jpg", ".png")
func FileExtension(value string) error {
	if !strings.HasPrefix(value, ".") {
		return errors.New("file extension must start with a dot")
	}
	if len(value) < 2 {
		return errors.New("file extension too short")
	}
	if !regexp.MustCompile(`^\.[a-zA-Z0-9]+$`).MatchString(value) {
		return errors.New("invalid file extension format")
	}
	return nil
}

// Validate URL format
func Urls(value string) error {
	u, err := url.ParseRequestURI(value)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return errors.New("invalid URL")
	}
	return nil
}

// Validate address: just check non-empty and reasonable length here (you can enhance)
func Address(value string) error {
	if len(strings.TrimSpace(value)) == 0 {
		return errors.New("address cannot be empty")
	}
	if len(value) > 200 {
		return errors.New("address too long")
	}
	return nil
}
