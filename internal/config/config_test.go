package config

import (
	"strings"
	"testing"
)

func TestLoadUsesDefaultsWhenEnvIsUnset(t *testing.T) {
	t.Setenv("APP_ENV", "")
	t.Setenv("PORT", "")
	t.Setenv("LOG_LEVEL", "")
	t.Setenv("SERVICE_NAME", "")

	cfg := Load()

	if cfg.AppEnv != "development" {
		t.Fatalf("expected default APP_ENV development, got %q", cfg.AppEnv)
	}

	if cfg.Port != "8080" {
		t.Fatalf("expected default PORT 8080, got %q", cfg.Port)
	}

	if cfg.LogLevel != "INFO" {
		t.Fatalf("expected default LOG_LEVEL INFO, got %q", cfg.LogLevel)
	}

	if cfg.ServiceName != "go-template" {
		t.Fatalf("expected default SERVICE_NAME go-template, got %q", cfg.ServiceName)
	}
}

func TestLoadUsesEnvironmentOverrides(t *testing.T) {
	t.Setenv("APP_ENV", "production")
	t.Setenv("PORT", "9090")
	t.Setenv("LOG_LEVEL", "DEBUG")
	t.Setenv("SERVICE_NAME", "acme-api")

	cfg := Load()

	if cfg.AppEnv != "production" {
		t.Fatalf("expected APP_ENV production, got %q", cfg.AppEnv)
	}

	if cfg.Port != "9090" {
		t.Fatalf("expected PORT 9090, got %q", cfg.Port)
	}

	if cfg.LogLevel != "DEBUG" {
		t.Fatalf("expected LOG_LEVEL DEBUG, got %q", cfg.LogLevel)
	}

	if cfg.ServiceName != "acme-api" {
		t.Fatalf("expected SERVICE_NAME acme-api, got %q", cfg.ServiceName)
	}
}

func TestLoadTrimsEnvironmentValues(t *testing.T) {
	t.Setenv("APP_ENV", " production ")
	t.Setenv("PORT", " 9090 ")
	t.Setenv("LOG_LEVEL", " debug ")
	t.Setenv("SERVICE_NAME", " acme-api ")

	cfg := Load()

	if cfg.AppEnv != "production" {
		t.Fatalf("expected trimmed APP_ENV production, got %q", cfg.AppEnv)
	}

	if cfg.Port != "9090" {
		t.Fatalf("expected trimmed PORT 9090, got %q", cfg.Port)
	}

	if cfg.LogLevel != "debug" {
		t.Fatalf("expected trimmed LOG_LEVEL debug, got %q", cfg.LogLevel)
	}

	if cfg.ServiceName != "acme-api" {
		t.Fatalf("expected trimmed SERVICE_NAME acme-api, got %q", cfg.ServiceName)
	}
}

func TestConfigValidateRejectsInvalidPort(t *testing.T) {
	cfg := Config{
		AppEnv:      "development",
		Port:        "70000",
		LogLevel:    "INFO",
		ServiceName: "go-template",
	}

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected invalid port to fail validation")
	}

	if !strings.Contains(err.Error(), "PORT") {
		t.Fatalf("expected validation error to mention PORT, got %v", err)
	}
}

func TestConfigValidateRejectsEmptyServiceName(t *testing.T) {
	cfg := Config{
		AppEnv:      "development",
		Port:        "8080",
		LogLevel:    "INFO",
		ServiceName: " ",
	}

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected empty service name to fail validation")
	}

	if !strings.Contains(err.Error(), "SERVICE_NAME") {
		t.Fatalf("expected validation error to mention SERVICE_NAME, got %v", err)
	}
}

func TestConfigValidateRejectsUnsupportedLogLevel(t *testing.T) {
	cfg := Config{
		AppEnv:      "development",
		Port:        "8080",
		LogLevel:    "verbose",
		ServiceName: "go-template",
	}

	err := cfg.Validate()
	if err == nil {
		t.Fatal("expected unsupported log level to fail validation")
	}

	if !strings.Contains(err.Error(), "LOG_LEVEL") {
		t.Fatalf("expected validation error to mention LOG_LEVEL, got %v", err)
	}
}

func TestConfigValidateAcceptsSupportedValues(t *testing.T) {
	cfg := Config{
		AppEnv:      "local",
		Port:        "8080",
		LogLevel:    "debug",
		ServiceName: "go-template",
	}

	if err := cfg.Validate(); err != nil {
		t.Fatalf("expected config to pass validation, got %v", err)
	}
}
