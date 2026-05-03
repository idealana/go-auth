package config

import (
	"go-auth/internal/helper"
)

func GetAppName() string {
	return helper.GetEnvString("APP_NAME", "Go App")
}

func GetAppPort() int {
	return helper.GetEnvInt("APP_PORT", 8000)
}
