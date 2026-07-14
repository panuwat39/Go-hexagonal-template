package userhttp

import (
	"github.com/gofiber/fiber/v3"

	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/command"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/usecase"
	apperror "github.com/panuwat39/go-hexagonal-template/internal/shared/errors"
	"github.com/panuwat39/go-hexagonal-template/internal/shared/response"
)

type UserHandler struct {
	createUserUseCase *usecase.CreateUserUseCase
}

func NewUserHandler(createUserUseCase *usecase.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	var request createUserRequest

	if err := c.Bind().JSON(&request); err != nil {
		return response.Error(c, apperror.New(
			apperror.CodeBadRequest,
			"invalid request body",
		))
	}

	output, err := h.createUserUseCase.Execute(c.Context(), command.CreateUserCommand{
		Email: request.Email,
		Name:  request.Name,
	})

	if err != nil {
		return response.Error(c, mapCreateUserError(err))
	}

	return response.Created(c, newCreateUserResponse(output))
}
