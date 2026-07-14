package middleware

import (
	"log/slog"

	"github.com/gofiber/fiber/v3"
	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

func NewRecoverMiddleware(logger *slog.Logger) fiber.Handler {
	return recoverer.New(recoverer.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c fiber.Ctx, err any) {
			logger.Error(
				"panic recovered",
				"error", err,
				"path", c.Path(),
				"method", c.Method(),
				"request_id", c.GetRespHeader(fiber.HeaderXRequestID),
			)
		},
	})
}
