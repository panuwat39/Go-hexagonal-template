package middleware

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v3"
)

func NewRequestLoggerMiddleware(logger *slog.Logger) fiber.Handler {
	return func(c fiber.Ctx) error {
		startedAt := time.Now()

		err := c.Next()

		latency := time.Since(startedAt)
		statusCode := c.Response().StatusCode()

		attrs := []any{
			"request_id", RequestID(c),
			"method", c.Method(),
			"path", c.Path(),
			"status", statusCode,
			"latency_ms", float64(latency.Microseconds()) / 1000.0,
			"ip", c.IP(),
			"user_agent", c.Get(fiber.HeaderUserAgent),
		}

		if err != nil {
			attrs = append(attrs, "error", err)
			logger.Error("http request failed", attrs...)

			return err
		}

		if statusCode >= fiber.StatusInternalServerError {
			logger.Error("http request completed", attrs...)
			return nil
		}

		if statusCode >= fiber.StatusBadRequest {
			logger.Warn("http request completed", attrs...)
			return nil
		}

		logger.Info("http request completed", attrs...)

		return nil
	}
}
