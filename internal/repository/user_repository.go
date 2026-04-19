package repository

import (
	"errors"
	"go-auth/internal/model/domain"
	"go-auth/internal/repository/iface"
)

func NewUserRepository() iface.UserRepositoryInterface {
	return &UserRepository{}
}

type UserRepository struct {
}

func (repository *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User

	// dummy data
	findUser := domain.User{
		ID: 1,
		FullName: "Uhuy Hehehe",
		Email: "uhuy@example.id",
		Password: "$2a$12$ibAELInQIpUUlM2jlH5KbOCRsTjO61KRPtF./1p0tFGBc/Y3q71Lq",
	}

	if email != findUser.Email {
		return user, errors.New("User not found.")
	}

	user = findUser

	return user, nil
}
