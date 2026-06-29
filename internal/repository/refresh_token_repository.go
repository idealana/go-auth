package repository

import (
	"context"
	
	"go-auth/internal/model/domain"

	"gorm.io/gorm"
)

type RefreshTokenRepository struct {
	*gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepositoryInterface {
	return &RefreshTokenRepository{
		DB: db,
	}
}

func (repo *RefreshTokenRepository) Insert(ctx context.Context, token *domain.RefreshToken) error {
	return repo.DB.
		WithContext(ctx).
		Create(token).
		Error
}
