package middleware

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/requestid"
)

func NewRequestIDMiddleware() fiber.Handler {
	return requestid.New(requestid.Config{
		Header:    fiber.HeaderXRequestID,
		Generator: newRequestID,
	})
}

func RequestID(c fiber.Ctx) string {
	requestID := c.GetRespHeader(fiber.HeaderXRequestID)
	if requestID != "" {
		return requestID
	}

	return c.Get(fiber.HeaderXRequestID)
}

func newRequestID() string {
	var b [16]byte

	if _, err := rand.Read(b[:]); err != nil {
		return "unknown"
	}

	return hex.EncodeToString(b[:])
}
