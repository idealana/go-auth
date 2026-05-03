package iface

import (
    "context"

	"go-auth/internal/dto"
    "go-auth/internal/security"
)

type AuthServiceInterface interface {
    Login(ctx context.Context, req dto.LoginRequest) (dto.LoginResult, error)
}

type TokenService interface {
    GenerateToken(security.UserClaims) (security.TokenPair, error)
}

type PasswordService interface {
    Verify(hashed, plain string) bool
}
