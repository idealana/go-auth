package database

import (
	"go-auth/internal/model/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
	)
	if err != nil {
		panic(err)
	}
}
