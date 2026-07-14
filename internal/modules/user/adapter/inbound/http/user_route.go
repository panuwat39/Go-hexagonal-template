package userhttp

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(router fiber.Router, handler *UserHandler) {
	users := router.Group("/users")

	users.Post("/", handler.CreateUser)
}
