package database

import (
	"go-auth/internal/model/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.User{},
		&domain.RefreshToken{},
	)
	if err != nil {
		panic(err)
	}
}
