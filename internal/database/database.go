package database

import (
	"time"

	"go-auth/internal/config"

	"gorm.io/driver/mysql"
    "gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDatabase() *gorm.DB {
	dbConfig := config.NewDatabaseConfig()

	dialect := mysql.Open(dbConfig.Username+":"+dbConfig.Password+"@tcp("+dbConfig.Host+":"+dbConfig.Port+")/"+dbConfig.Database+"?charset=utf8mb4&parseTime=True&loc=Local")
    db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConns)
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Duration(dbConfig.ConnMaxLifetime) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(dbConfig.ConnMaxIdleTime) * time.Minute)

	return db
}
