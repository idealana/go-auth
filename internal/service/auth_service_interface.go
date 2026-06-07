package service

import (
    "context"

	"go-auth/internal/model/dto"
    "go-auth/internal/security"
)

type AuthServiceInterface interface {
    Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResult, error)
    GetProfile(ctx context.Context, userId uint) (dto.ProfileResult, error)
}

type TokenService interface {
    GenerateToken(security.UserClaims) (security.TokenPair, error)
}

type PasswordService interface {
    Verify(hashed, plain string) error
}
