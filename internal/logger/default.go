package logger

import (
	"log/slog"
)

type defaultLogger struct {}

func (defaultLogger) Fatal(message string, args ...any) {
	slog.Error(message, args...)
	ExitApp(1)
}

func (defaultLogger) Error(message string, args ...any) {
	slog.Error(message, args...)
}

func (defaultLogger) Warn(message string, args ...any) {
	slog.Warn(message, args...)
}

func (defaultLogger) Info(message string, args ...any) {
	slog.Info(message, args...)
}

func (defaultLogger) Debug(message string, args ...any) {
	slog.Debug(message, args...)
}
