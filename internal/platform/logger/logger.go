package logger

import (
	"log/slog"
	"os"
	"strings"
)

type Config struct {
	Env string
}

func New(config Config) *slog.Logger {
	level := slog.LevelInfo

	if strings.EqualFold(config.Env, "local") || strings.EqualFold(config.Env, "development") {
		level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	return slog.New(handler)
}
