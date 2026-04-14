package observability

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(level string) *zap.Logger {
	logLevel := zap.NewAtomicLevel()

	switch strings.ToUpper(level) {
	case "DEBUG":
		logLevel.SetLevel(zap.DebugLevel)
	case "WARN":
		logLevel.SetLevel(zap.WarnLevel)
	case "ERROR":
		logLevel.SetLevel(zap.ErrorLevel)
	default:
		logLevel.SetLevel(zap.InfoLevel)
	}

	cfg := zap.Config{
		Level:            logLevel,
		Development:      false,
		Encoding:         "json",
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.MillisDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
	}

	return zap.Must(cfg.Build())
}
