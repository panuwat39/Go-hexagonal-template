package validator

import (
	"strings"
)

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	fieldErrors []FieldError
}

func NewValidationError() *ValidationError {
	return &ValidationError{
		fieldErrors: make([]FieldError, 0),
	}
}

func (e *ValidationError) Error() string {
	return "validation failed"
}

func (e *ValidationError) Add(field string, message string) {
	e.fieldErrors = append(e.fieldErrors, FieldError{
		Field:   field,
		Message: message,
	})
}

func (e *ValidationError) HasErrors() bool {
	return len(e.fieldErrors) > 0
}

func (e *ValidationError) Fields() []FieldError {
	return e.fieldErrors
}

func Required(errors *ValidationError, field string, value string) {
	if strings.TrimSpace(value) == "" {
		errors.Add(field, field+" is required")
	}
}

func MinLength(errors *ValidationError, field string, value string, min int) {
	if strings.TrimSpace(value) == "" {
		return
	}

	if len(strings.TrimSpace(value)) < min {
		errors.Add(field, field+" must be at least "+itoa(min)+" characters")
	}
}

func Email(errors *ValidationError, field string, value string) {
	normalizedValue := strings.TrimSpace(value)

	if normalizedValue == "" {
		return
	}

	if !strings.Contains(normalizedValue, "@") {
		errors.Add(field, field+" must be a valid email")
	}
}

func itoa(value int) string {
	if value == 0 {
		return "0"
	}

	digits := make([]byte, 0)

	for value > 0 {
		digits = append([]byte{byte('0' + value%10)}, digits...)
		value /= 10
	}

	return string(digits)
}
