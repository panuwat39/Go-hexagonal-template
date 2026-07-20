package apperror

type Code string

const (
	CodeBadRequest   Code = "BAD_REQUEST"
	CodeUnauthorized Code = "UNAUTHORIZED"
	CodeForbidden    Code = "FORBIDDEN"
	CodeNotFound     Code = "NOT_FOUND"
	CodeConflict     Code = "CONFLICT"
	CodeInternal     Code = "INTERNAL"
)

type AppError struct {
	code    Code
	message string
	cause   error
	details any
}

func New(code Code, message string) *AppError {
	return &AppError{
		code:    code,
		message: message,
	}
}

func NewWithDetails(code Code, message string, details any) *AppError {
	return &AppError{
		code:    code,
		message: message,
		details: details,
	}
}

func Wrap(code Code, message string, cause error) *AppError {
	return &AppError{
		code:    code,
		message: message,
		cause:   cause,
	}
}

func WrapWithDetails(code Code, message string, cause error, details any) *AppError {
	return &AppError{
		code:    code,
		message: message,
		cause:   cause,
		details: details,
	}
}

func (e *AppError) Error() string {
	return e.message
}

func (e *AppError) Unwrap() error {
	return e.cause
}

func (e *AppError) Code() Code {
	return e.code
}

func (e *AppError) Message() string {
	return e.message
}

func (e *AppError) Details() any {
	return e.details
}

func (e *AppError) HasDetails() bool {
	return e.details != nil
}
