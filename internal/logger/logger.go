package logger

import (
	"os"
)

var ExitApp = os.Exit

type Logger interface {
	Fatal(message string, args ...any)
	Error(message string, args ...any)
    Warn(message string, args ...any)
    Info(message string, args ...any)
    Debug(message string, args ...any)
}

func Resolve(loggers ...Logger) Logger {
	if len(loggers) > 0 && loggers[0] != nil {
		return loggers[0]
	}
	return defaultLogger{}
}
