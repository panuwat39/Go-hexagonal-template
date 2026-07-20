package userhttp

import (
	apperror "github.com/panuwat39/go-hexagonal-template/internal/shared/errors"
	"github.com/panuwat39/go-hexagonal-template/internal/shared/validator"
)

type createUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (r createUserRequest) Validate() error {
	validationError := validator.NewValidationError()

	validator.Required(validationError, "email", r.Email)
	validator.Email(validationError, "email", r.Email)

	validator.Required(validationError, "name", r.Name)
	validator.MinLength(validationError, "name", r.Name, 2)

	if validationError.HasErrors() {
		return apperror.NewWithDetails(
			apperror.CodeBadRequest,
			"validation failed",
			validationError.Fields(),
		)
	}

	return nil
}
