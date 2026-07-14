package user

import (
	"github.com/gofiber/fiber/v3"

	userhttp "github.com/panuwat39/go-hexagonal-template/internal/modules/user/adapter/inbound/http"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/adapter/outbound/persistence"
	"github.com/panuwat39/go-hexagonal-template/internal/modules/user/application/usecase"
)

type Module struct {
	handler *userhttp.UserHandler
}

func NewModule() *Module {
	userRepository := persistence.NewInMemoryUserRepository()
	createUserUseCase := usecase.NewCreateUserUseCase(userRepository)
	handler := userhttp.NewUserHandler(createUserUseCase)

	return &Module{
		handler: handler,
	}
}

func (m *Module) RegisterRoutes(app *fiber.App) {
	m.handler.RegisterRoutes(app)
}
