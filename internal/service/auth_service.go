package service

import (
    "context"
    "errors"
    "fmt"
    "time"

    "go-auth/internal/apperror"
    "go-auth/internal/config"
    "go-auth/internal/helper"
    "go-auth/internal/model/domain"
	"go-auth/internal/model/dto"
    "go-auth/internal/repository"
)

func NewAuthService(
    userRepository repository.UserRepositoryInterface,
    refrehTokenRepository repository.RefreshTokenRepositoryInterface,
    tokenService TokenService,
    passwordService PasswordService,
) *AuthService {
    return &AuthService{
        UserRepository: userRepository,
        RefreshTokenRepository: refrehTokenRepository,
        TokenService: tokenService,
		PasswordService: passwordService,
    }
}

type AuthService struct {
    UserRepository repository.UserRepositoryInterface
    RefreshTokenRepository repository.RefreshTokenRepositoryInterface
    TokenService TokenService
	PasswordService PasswordService
}

func (svc *AuthService) Login(ctx context.Context, req dto.LoginRequest, reqInfo *dto.RequestInfo) (dto.LoginResult, error) {
	user, err := svc.UserRepository.FindByEmail(ctx, req.Email)
    if err != nil {
        if errors.Is(err, apperror.ErrNotFound) {
            return dto.LoginResult{}, apperror.ErrInvalidCredentials
        }

        return dto.LoginResult{}, fmt.Errorf("find user: %w", err)
    }
	
    if err := svc.PasswordService.Verify(user.Password, req.Password); err != nil {
        return dto.LoginResult{}, apperror.ErrInvalidCredentials
    }

    tokenPair, err := svc.TokenService.GenerateToken(&user)
    if err != nil {
        return dto.LoginResult{}, err
    }

    refreshToken := domain.RefreshToken{
        UserID: user.ID,
        TokenHash: helper.HashToken(tokenPair.RefreshToken),
        ExpiredAt: time.Now().Add(time.Duration(config.GetJWTRefreshExpired()) * 24 * time.Hour),
        IPAddress: &reqInfo.IPAddress,
        UserAgent: &reqInfo.UserAgent,
    }

    if err := svc.RefreshTokenRepository.Insert(ctx, &refreshToken); err != nil {
        return dto.LoginResult{}, apperror.ErrCreateRefreshToken
    }

    return dto.LoginResult{
		UserID: user.ID,
		AccessToken: tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (svc *AuthService) GetProfile(ctx context.Context, userId uint) (dto.ProfileResult, error) {
    user, err := svc.UserRepository.FindById(ctx, userId)
    if err != nil {
        if errors.Is(err, apperror.ErrNotFound) {
            return dto.ProfileResult{}, apperror.ErrNotFound
        }

        return dto.ProfileResult{}, fmt.Errorf("find user: %w", err)
    }

    return dto.ProfileResult{
		UserID: user.ID,
		Email: user.Email,
        Role: user.Role,
        Status: user.Status,
	}, nil
}
