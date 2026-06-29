package repository

import (
	"context"
	
	"go-auth/internal/model/domain"
)

type RefreshTokenRepositoryInterface interface {
	Insert(ctx context.Context, token *domain.RefreshToken) error
}
