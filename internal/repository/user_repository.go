package repository

import (
	"context"
	
	"go-auth/internal/apperror"
	"go-auth/internal/model/domain"
)

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

type UserRepository struct {
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	// dummy data
	findUser := domain.User{
		ID: 1,
		FullName: "Uhuy Hehehe",
		Email: "uhuy@example.id",
		Password: "$2a$12$ibAELInQIpUUlM2jlH5KbOCRsTjO61KRPtF./1p0tFGBc/Y3q71Lq",
	}

	if email != findUser.Email {
		return user, apperror.ErrNotFound
	}

	user = findUser

	return user, nil
}
