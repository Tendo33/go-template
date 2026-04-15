package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	AppEnv      string
	Port        string
	LogLevel    string
	ServiceName string
}

func Load() Config {
	return Config{
		AppEnv:      getEnv("APP_ENV", "development"),
		Port:        getEnv("PORT", "8080"),
		LogLevel:    getEnv("LOG_LEVEL", "INFO"),
		ServiceName: getEnv("SERVICE_NAME", "go-template"),
	}
}

func (c Config) Validate() error {
	port, err := strconv.Atoi(strings.TrimSpace(c.Port))
	if err != nil || port < 1 || port > 65535 {
		return fmt.Errorf("invalid PORT %q: must be a number between 1 and 65535", c.Port)
	}

	if strings.TrimSpace(c.ServiceName) == "" {
		return fmt.Errorf("invalid SERVICE_NAME: must not be empty")
	}

	switch strings.ToUpper(strings.TrimSpace(c.LogLevel)) {
	case "DEBUG", "INFO", "WARN", "ERROR":
		return nil
	default:
		return fmt.Errorf("invalid LOG_LEVEL %q: supported values are DEBUG, INFO, WARN, ERROR", c.LogLevel)
	}
}

func getEnv(key string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback
	}

	return value
}
