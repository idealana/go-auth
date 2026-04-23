package service

import (
	"errors"
    "go-auth/internal/model/domain"
	"go-auth/internal/model/web"
    "go-auth/internal/repository/iface"
    "go-auth/pkg/utils"
)

func NewAuthService(userRepository iface.UserRepositoryInterface, tokenUtility *utils.TokenUtility) *AuthService {
    return &AuthService{
        UserRepository: userRepository,
        TokenUtility: tokenUtility,
    }
}

type AuthService struct {
    UserRepository iface.UserRepositoryInterface
    TokenUtility *utils.TokenUtility
}

func (service *AuthService) Login(req web.LoginRequest) (string, string, error) {
	_, err := service.UserRepository.FindByEmail(req.Email)

    if err != nil {
        return res, "", errors.New("Invalid email or password.")
    }

    isVerify := utils.VerifyPassword(user.Password, req.Password)

    if !isVerify {
        return res, "", errors.New("Invalid email or password.")
    }

    accessToken, refreshToken, err := service.TokenUtility.GenerateToken(&domain.User{ID: user.ID})

    if err != nil {
        return res, "", errors.New("Failed to generate token.")
    }

    return accessToken, refreshToken, nil
}
