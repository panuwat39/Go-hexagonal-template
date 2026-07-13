package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := fiber.New(fiber.Config{
		AppName: "go-hexagonal-template",
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})

	go func() {
		logger.Info("http server started", "addr", ":8080")

		if err := app.Listen(":8080"); err != nil {
			logger.Error("http server failed", "error", err)
			os.Exit(1)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	if err := app.Shutdown(); err != nil {
		logger.Error("http server shutdown failed", "error", err)
		os.Exit(1)
	}

	logger.Info("http server stopped")
}
