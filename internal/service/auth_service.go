package service

import (
	"errors"
	"go-auth/internal/model/web"
    "go-auth/internal/repository/iface"
    "go-auth/pkg/utils"
)

func NewAuthService(userRepository iface.UserRepositoryInterface) *AuthService {
    return &AuthService{
        UserRepository: userRepository,
    }
}

type AuthService struct {
    UserRepository iface.UserRepositoryInterface
}

func (service *AuthService) Login(req web.LoginRequest) (web.LoginResponse, error) {
	user, err := service.UserRepository.FindByEmail(req.Email)
    res := web.NewLoginResponse(user)

    if err != nil {
        return res, errors.New("Invalid email or password.")
    }

    isVerify := utils.VerifyPassword(user.Password, req.Password)

    if !isVerify {
        return res, errors.New("Invalid email or password.")
    }

    return res, nil
}
