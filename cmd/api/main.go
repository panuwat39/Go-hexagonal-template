package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v3"

	"github.com/panuwat39/go-hexagonal-template/internal/bootstrap"
	usermodule "github.com/panuwat39/go-hexagonal-template/internal/modules/user"
)

func main() {
	cfg := bootstrap.LoadConfig()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := fiber.New(fiber.Config{
		AppName: cfg.App.Name,
	})

	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"env":    cfg.App.Env,
		})
	})

	userModule := usermodule.NewModule()
	userModule.RegisterRoutes(app)

	go func() {
		logger.Info("http server started", "addr", cfg.HTTPAddress())

		if err := app.Listen(cfg.HTTPAddress()); err != nil {
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
