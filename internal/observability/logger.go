package observability

import (
	"context"
	"os"
	"strings"
	"sync/atomic"

	"github.com/Tendo33/go-template/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextKey string

const (
	loggerContextKey    contextKey = "logger"
	requestIDContextKey contextKey = "request_id"
	traceIDContextKey   contextKey = "trace_id"
	spanIDContextKey    contextKey = "span_id"
)

var fallbackLogger atomic.Value

func init() {
	fallbackLogger.Store(zap.NewNop())
}

func NewLogger(cfg config.Config) *zap.Logger {
	return newLogger(cfg, zapcore.AddSync(os.Stdout), zapcore.AddSync(os.Stderr))
}

func newLogger(cfg config.Config, output, errOutput zapcore.WriteSyncer) *zap.Logger {
	logLevel := zap.NewAtomicLevel()

	switch strings.ToUpper(cfg.LogLevel) {
	case "DEBUG":
		logLevel.SetLevel(zap.DebugLevel)
	case "WARN":
		logLevel.SetLevel(zap.WarnLevel)
	case "ERROR":
		logLevel.SetLevel(zap.ErrorLevel)
	default:
		logLevel.SetLevel(zap.InfoLevel)
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	if IsDevelopmentEnv(cfg.AppEnv) {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	logger := zap.New(
		zapcore.NewCore(encoder, output, logLevel),
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
		zap.ErrorOutput(errOutput),
		zap.Fields(
			zap.String("service", cfg.ServiceName),
			zap.String("env", cfg.AppEnv),
		),
	)

	setFallbackLogger(logger)

	return logger
}

func IsDevelopmentEnv(appEnv string) bool {
	switch strings.ToLower(strings.TrimSpace(appEnv)) {
	case "development", "dev", "local":
		return true
	default:
		return false
	}
}

func WithLogger(ctx context.Context, logger *zap.Logger) context.Context {
	if logger == nil {
		return ctx
	}

	return context.WithValue(ctx, loggerContextKey, logger)
}

func FromContext(ctx context.Context) *zap.Logger {
	logger := fallbackLogger.Load().(*zap.Logger)

	if ctx != nil {
		if contextLogger, ok := ctx.Value(loggerContextKey).(*zap.Logger); ok && contextLogger != nil {
			logger = contextLogger
		}
	}

	fields := ContextFields(ctx)
	if len(fields) == 0 {
		return logger
	}

	return logger.With(fields...)
}

func WithRequestID(ctx context.Context, requestID string) context.Context {
	if requestID == "" {
		return ctx
	}

	return context.WithValue(ctx, requestIDContextKey, requestID)
}

func RequestIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	requestID, _ := ctx.Value(requestIDContextKey).(string)
	return requestID
}

func WithTraceContext(ctx context.Context, traceID, spanID string) context.Context {
	if traceID != "" {
		ctx = context.WithValue(ctx, traceIDContextKey, traceID)
	}

	if spanID != "" {
		ctx = context.WithValue(ctx, spanIDContextKey, spanID)
	}

	return ctx
}

func TraceIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	traceID, _ := ctx.Value(traceIDContextKey).(string)
	return traceID
}

func SpanIDFromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}

	spanID, _ := ctx.Value(spanIDContextKey).(string)
	return spanID
}

func ContextFields(ctx context.Context) []zap.Field {
	fields := make([]zap.Field, 0, 3)

	if requestID := RequestIDFromContext(ctx); requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}

	if traceID := TraceIDFromContext(ctx); traceID != "" {
		fields = append(fields, zap.String("trace_id", traceID))
	}

	if spanID := SpanIDFromContext(ctx); spanID != "" {
		fields = append(fields, zap.String("span_id", spanID))
	}

	return fields
}

func setFallbackLogger(logger *zap.Logger) {
	if logger == nil {
		logger = zap.NewNop()
	}

	fallbackLogger.Store(logger)
}
