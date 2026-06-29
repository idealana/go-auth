package config

import (
	"errors"

	"go-auth/internal/helper"
)

func GetJWTAccessKey() (string, error) {
	key := helper.GetEnvString("JWT_ACCESS_KEY")
	if key == "" {
		return "", errors.New("JWT_ACCESS_KEY environment variable is required")
	}

	return key, nil
}

func GetJWTAccessExpired() int {
	defaultExpired := 15
	return helper.GetEnvInt("JWT_ACCESS_EXPIRED", defaultExpired)
}

func GetJWTRefreshExpired() int {
	defaultExpired := 7
	return helper.GetEnvInt("JWT_REFRESH_EXPIRED", defaultExpired)
}
