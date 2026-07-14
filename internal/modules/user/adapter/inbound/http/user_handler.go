package userhttp

import (
	"errors"

	"github.com/gofiber/fiber/v3"

	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/command"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/usecase"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/domain/entity"
)

type UserHandler struct {
	createUserUseCase *usecase.CreateUserUseCase
}

func NewUserHandler(createUserUseCase *usecase.CreateUserUseCase) *UserHandler {
	return &UserHandler{
		createUserUseCase: createUserUseCase,
	}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	app.Post("/users", h.CreateUser)
}

func (h *UserHandler) CreateUser(c fiber.Ctx) error {
	var request createUserRequest

	if err := c.Bind().JSON(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: "invalid request body",
		})
	}

	output, err := h.createUserUseCase.Execute(c.Context(), command.CreateUserCommand{
		Email: request.Email,
		Name:  request.Name,
	})

	if err != nil {
		return handleCreateUserError(c, err)
	}

	return c.Status(fiber.StatusCreated).JSON(createUserResponse{
		ID:        output.ID,
		Email:     output.Email,
		Name:      output.Name,
		CreatedAt: output.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
	})
}

func handleCreateUserError(c fiber.Ctx, err error) error {
	switch {
	case errors.Is(err, entity.ErrInvalidEmail), errors.Is(err, entity.ErrInvalidName):
		return c.Status(fiber.StatusBadRequest).JSON(errorResponse{
			Message: err.Error(),
		})

	case errors.Is(err, usecase.ErrUserAlreadyExists):
		return c.Status(fiber.StatusConflict).JSON(errorResponse{
			Message: "email already exists",
		})

	default:
		return c.Status(fiber.StatusInternalServerError).JSON(errorResponse{
			Message: "internal server error",
		})
	}
}

type createUserRequest struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type createUserResponse struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type errorResponse struct {
	Message string `json:"message"`
}
