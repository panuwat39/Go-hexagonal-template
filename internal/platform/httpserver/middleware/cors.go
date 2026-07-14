package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func NewCORSMiddleware(allowedOrigins string) fiber.Handler {
	origins := parseAllowedOrigins(allowedOrigins)

	return cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodPut,
			fiber.MethodPatch,
			fiber.MethodDelete,
			fiber.MethodOptions,
		},
		AllowHeaders: []string{
			fiber.HeaderOrigin,
			fiber.HeaderContentType,
			fiber.HeaderAccept,
			fiber.HeaderAuthorization,
			fiber.HeaderXRequestID,
		},
		ExposeHeaders: []string{
			fiber.HeaderXRequestID,
		},
	})
}

func parseAllowedOrigins(value string) []string {
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return []string{"*"}
	}

	parts := strings.Split(trimmedValue, ",")
	origins := make([]string, 0, len(parts))

	for _, part := range parts {
		origin := strings.TrimSpace(part)
		if origin != "" {
			origins = append(origins, origin)
		}
	}

	if len(origins) == 0 {
		return []string{"*"}
	}

	return origins
}
