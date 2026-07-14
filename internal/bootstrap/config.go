package bootstrap

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	App  AppConfig
	HTTP HTTPConfig
}

type AppConfig struct {
	Name string
	Env  string
}

type HTTPConfig struct {
	Port int
}

func LoadConfig() Config {
	return Config{
		App: AppConfig{
			Name: getEnv("APP_NAME", "go-hexagonal-template"),
			Env:  getEnv("APP_ENV", "local"),
		},
		HTTP: HTTPConfig{
			Port: getEnvAsInt("HTTP_PORT", 8080),
		},
	}
}

func (c Config) HTTPAddress() string {
	return fmt.Sprintf(":%d", c.HTTP.Port)
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsedValue
}
