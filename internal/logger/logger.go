package logger

type Logger interface {
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
