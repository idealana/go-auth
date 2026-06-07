package repository

import (
	"context"
	
	"go-auth/internal/model/domain"
)

type UserRepositoryInterface interface {
	FindByEmail(ctx context.Context, email string) (domain.User, error)
	FindById(ctx context.Context, id uint) (domain.User, error)
}
