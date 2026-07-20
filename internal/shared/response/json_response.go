package response

import (
	"errors"

	"github.com/gofiber/fiber/v3"

	apperror "github.com/panuwat39/go-hexagonal-template/internal/shared/errors"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func JSON(c fiber.Ctx, statusCode int, body any) error {
	return c.Status(statusCode).JSON(body)
}

func OK(c fiber.Ctx, body any) error {
	return JSON(c, StatusOK, body)
}

func Created(c fiber.Ctx, body any) error {
	return JSON(c, StatusCreated, body)
}

func NoContent(c fiber.Ctx) error {
	return c.SendStatus(StatusNoContent)
}

func Error(c fiber.Ctx, err error) error {
	var appErr *apperror.AppError

	if errors.As(err, &appErr) {
		return JSON(c, statusFromCode(appErr.Code()), ErrorResponse{
			Code:    string(appErr.Code()),
			Message: appErr.Message(),
			Details: appErr.Details(),
		})
	}

	return JSON(c, StatusInternalServerError, ErrorResponse{
		Code:    string(apperror.CodeInternal),
		Message: "internal server error",
	})
}

func statusFromCode(code apperror.Code) int {
	switch code {
	case apperror.CodeBadRequest:
		return StatusBadRequest
	case apperror.CodeUnauthorized:
		return StatusUnauthorized
	case apperror.CodeForbidden:
		return StatusForbidden
	case apperror.CodeNotFound:
		return StatusNotFound
	case apperror.CodeConflict:
		return StatusConflict
	default:
		return StatusInternalServerError
	}
}
