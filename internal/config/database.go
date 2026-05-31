package config

import (
	"go-auth/internal/helper"
)

type DatabaseConfig struct {
	Host string
	Port string
	Database string
	Username string
	Password string
	MaxOpenConns int
	MaxIdleConns int
	ConnMaxLifetime int
	ConnMaxIdleTime int
}

func NewDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Host: helper.GetEnvString("DB_HOST"),
		Port: helper.GetEnvString("DB_PORT"),
		Database: helper.GetEnvString("DB_DATABASE"),
		Username: helper.GetEnvString("DB_USERNAME"),
		Password: helper.GetEnvString("DB_PASSWORD"),
		MaxOpenConns: helper.GetEnvInt("DB_MAX_OPEN_CONNS", 100),
		MaxIdleConns: helper.GetEnvInt("DB_MAX_IDLE_CONNS", 10),
		ConnMaxLifetime: helper.GetEnvInt("DB_CONN_MAX_LIFETIME", 30),
		ConnMaxIdleTime: helper.GetEnvInt("DB_CONN_MAX_IDLE_TIME", 5),
	}
}
