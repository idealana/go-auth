package helper

import (
	"os"
	"strconv"
)

func GetEnvString(key string, defaultValue ...string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return ""
}

func GetEnvInt(key string, defaultValue ...int) int {
	if value, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return value
	}

	if len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return 0
}
