package middleware

import (
    "context"
    
    "go-auth/internal/helper"
    "go-auth/internal/http/response"
    "go-auth/internal/logger"
    "go-auth/internal/model/domain"
  
    "github.com/gofiber/fiber/v3"
)

type authKey struct {
	name string
}

var defaultAuthKey = authKey{"validated_auth"}

type TokenParser interface {
    ParseAccessToken(ctx context.Context, token string) (*domain.Auth, error)
}

type Auth struct {
    tokenParser TokenParser
    log logger.Logger
}

func NewAuth(tokenParser TokenParser, log logger.Logger) *Auth {
    return &Auth{
        tokenParser: tokenParser,
        log: log,
    }
}

func (auth *Auth) ValidateAuth() fiber.Handler {
    return func(ctx fiber.Ctx) error {
        authHeader := ctx.Get("Authorization")
        if authHeader == "" {
            auth.log.Warn("missing authorization header", "path", ctx.Path())

            return response.Unauthorized(ctx, "Missing Authorization Header")
        }

        token := helper.ParseBearerToken(authHeader)
        if token == "" {
            auth.log.Warn("invalid authorization format", "path", ctx.Path())

            return response.Unauthorized(ctx, "Invalid Authorization Format")
        }
        
        authUser, err := auth.tokenParser.ParseAccessToken(ctx.Context(), token)
        if err != nil {
            auth.log.Error("failed to parse token", "error", err, "path", ctx.Path())

            return response.Unauthorized(ctx, "Invalid Access Token")
        }

        ctx.Locals(defaultAuthKey, authUser)
        return ctx.Next()
    }
}

func GetAuth(ctx fiber.Ctx) (*domain.Auth, bool) {
    val := ctx.Locals(defaultAuthKey)
    if val == nil {
        return nil, false
    }
    
    auth, ok := val.(*domain.Auth)
    return auth, ok
}
