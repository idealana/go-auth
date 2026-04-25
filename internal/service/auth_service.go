package service

import (
	"errors"
    "go-auth/internal/model/domain"
	"go-auth/internal/model/dto"
    "go-auth/internal/repository/iface"
)

var (
	ErrInvalidCredentials = errors.New("invalid email or password")
)

type TokenService interface {
    GenerateToken(*domain.User) (string, string, error)
}

type PasswordService interface {
    Verify(hashed, plain string) bool
}

func NewAuthService(userRepository iface.UserRepositoryInterface, tokenService TokenService, passwordService PasswordService) *AuthService {
    return &AuthService{
        UserRepository: userRepository,
        TokenService: tokenService,
		PasswordService: passwordService,
    }
}

type AuthService struct {
    UserRepository iface.UserRepositoryInterface
    TokenService TokenService
	PasswordService PasswordService
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := s.UserRepository.FindByEmail(req.Email)
    if err != nil {
        return dto.LoginResponse{}, ErrInvalidCredentials
    }
	
    if !s.PasswordService.Verify(user.Password, req.Password) {
        return dto.LoginResponse{}, ErrInvalidCredentials
    }

    accessToken, refreshToken, err := s.TokenService.GenerateToken(&user)
    if err != nil {
        return dto.LoginResponse{}, err
    }

    return dto.LoginResponse{
		UserID: user.ID,
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}
