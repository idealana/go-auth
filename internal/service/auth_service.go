package service

import (
    "context"
    "errors"
    "fmt"

    "go-auth/internal/apperror"
	"go-auth/internal/dto"
    "go-auth/internal/repository"
)

func NewAuthService(userRepository repository.UserRepositoryInterface, tokenService TokenService, passwordService PasswordService) *AuthService {
    return &AuthService{
        UserRepository: userRepository,
        TokenService: tokenService,
		PasswordService: passwordService,
    }
}

type AuthService struct {
    UserRepository repository.UserRepositoryInterface
    TokenService TokenService
	PasswordService PasswordService
}

func (svc *AuthService) Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResult, error) {
	user, err := svc.UserRepository.FindByEmail(ctx, req.Email)
    if err != nil {
        if errors.Is(err, apperror.ErrNotFound) {
            return dto.LoginResult{}, apperror.ErrInvalidCredentials
        }

        return dto.LoginResult{}, fmt.Errorf("find user: %w", err)
    }
	
    if !svc.PasswordService.Verify(user.Password, req.Password) {
        return dto.LoginResult{}, apperror.ErrInvalidCredentials
    }

    tokenPair, err := svc.TokenService.GenerateToken(&user)
    if err != nil {
        return dto.LoginResult{}, err
    }

    return dto.LoginResult{
		UserID: user.ID,
		AccessToken: tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}
