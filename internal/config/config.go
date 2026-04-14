package config

import "os"

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

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
