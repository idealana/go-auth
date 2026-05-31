package repository

import (
	"context"
	"errors"
	
	"go-auth/internal/apperror"
	"go-auth/internal/model/domain"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &UserRepository{
		DB: db,
	}
}

type UserRepository struct {
	*gorm.DB
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	err := repo.DB.
		WithContext(ctx).
		Where(&domain.User{Email: email}).
		First(&user).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, apperror.ErrNotFound
		}
		return domain.User{}, err
	}

	return user, nil
}
