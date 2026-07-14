package userhttp

import (
	"errors"

	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/usecase"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/domain/entity"
	apperror "github.com/panuwat39/go-hexagonal-template/internal/shared/errors"
)

func mapCreateUserError(err error) error {
	switch {
	case errors.Is(err, entity.ErrInvalidEmail), errors.Is(err, entity.ErrInvalidName):
		return apperror.New(apperror.CodeBadRequest, err.Error())

	case errors.Is(err, usecase.ErrUserAlreadyExists):
		return apperror.New(apperror.CodeConflict, "email already exists")

	default:
		return apperror.Wrap(apperror.CodeInternal, "internal server error", err)
	}
}
