package service

import (
	"errors"
    "go-auth/internal/model/domain"
	"go-auth/internal/model/web"
    "go-auth/internal/repository/iface"
    "go-auth/pkg/utils"
)

func NewAuthService(userRepository iface.UserRepositoryInterface, jwt *utils.JWT) *AuthService {
    return &AuthService{
        UserRepository: userRepository,
        JWT: jwt,
    }
}

type AuthService struct {
    UserRepository iface.UserRepositoryInterface
    JWT *utils.JWT
}

func (service *AuthService) Login(req web.LoginRequest) (web.LoginResponse, string, error) {
	user, err := service.UserRepository.FindByEmail(req.Email)
    res := web.NewLoginResponse(user)

    if err != nil {
        return res, "", errors.New("Invalid email or password.")
    }

    isVerify := utils.VerifyPassword(user.Password, req.Password)

    if !isVerify {
        return res, "", errors.New("Invalid email or password.")
    }

    token, err := service.JWT.GenerateToken(&domain.User{ID: user.ID})

    if err != nil {
        return res, "", errors.New("Failed to generate token.")
    }

    return res, token, nil
}
