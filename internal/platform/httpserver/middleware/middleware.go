package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type Config struct {
	AllowedOrigins string
}

func Register(app *fiber.App, logger *slog.Logger, config Config) {
	app.Use(NewRecoverMiddleware(logger))
	app.Use(NewRequestIDMiddleware())
	app.Use(NewRequestLoggerMiddleware(logger))
	app.Use(NewCORSMiddleware(config.AllowedOrigins))
}
