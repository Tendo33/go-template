package observability

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Tendo33/go-template/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLoggerUsesColorConsoleEncoderInDevelopment(t *testing.T) {
	var output bytes.Buffer

	logger := newLogger(
		config.Config{
			AppEnv:      "development",
			LogLevel:    "INFO",
			ServiceName: "go-template",
		},
		zapcore.AddSync(&output),
		zapcore.AddSync(&output),
	)

	logger.Info("development log")

	line := output.String()
	if !strings.Contains(line, "development log") {
		t.Fatalf("expected console output to contain message, got %q", line)
	}

	if !strings.Contains(line, "\x1b[") {
		t.Fatalf("expected development console output to contain ANSI color codes, got %q", line)
	}

	if strings.HasPrefix(strings.TrimSpace(line), "{") {
		t.Fatalf("expected development output to be console formatted, got JSON: %q", line)
	}
}

func TestNewLoggerUsesJSONEncoderOutsideDevelopment(t *testing.T) {
	var output bytes.Buffer

	logger := newLogger(
		config.Config{
			AppEnv:      "production",
			LogLevel:    "INFO",
			ServiceName: "go-template",
		},
		zapcore.AddSync(&output),
		zapcore.AddSync(&output),
	)

	logger.Info("production log")

	line := strings.TrimSpace(output.String())
	if !strings.HasPrefix(line, "{") {
		t.Fatalf("expected production output to be JSON, got %q", line)
	}

	if !strings.Contains(line, "\"service\":\"go-template\"") {
		t.Fatalf("expected production output to include service field, got %q", line)
	}

	if !strings.Contains(line, "\"env\":\"production\"") {
		t.Fatalf("expected production output to include env field, got %q", line)
	}
}

func TestFromContextFallsBackToBaseLogger(t *testing.T) {
	baseLogger := zap.NewNop()
	setFallbackLogger(baseLogger)

	logger := FromContext(t.Context())
	if logger != baseLogger {
		t.Fatalf("expected FromContext to return fallback logger when context has no logger")
	}
}

func TestRequestIDFromContextReturnsStoredValue(t *testing.T) {
	ctx := WithRequestID(t.Context(), "test-123")

	if got := RequestIDFromContext(ctx); got != "test-123" {
		t.Fatalf("expected request ID test-123, got %q", got)
	}
}

func TestContextFieldsIncludeTraceIdentifiers(t *testing.T) {
	ctx := WithRequestID(t.Context(), "req-123")
	ctx = WithTraceContext(ctx, "trace-123", "span-456")

	fields := ContextFields(ctx)

	observed := map[string]string{}
	for _, field := range fields {
		observed[field.Key] = field.String
	}

	if observed["request_id"] != "req-123" {
		t.Fatalf("expected request_id req-123, got %q", observed["request_id"])
	}

	if observed["trace_id"] != "trace-123" {
		t.Fatalf("expected trace_id trace-123, got %q", observed["trace_id"])
	}

	if observed["span_id"] != "span-456" {
		t.Fatalf("expected span_id span-456, got %q", observed["span_id"])
	}
}
