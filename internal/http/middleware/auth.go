package middleware

import (
    "context"
    "log"
    "strings"
    
    "go-auth/internal/model/domain"
  
    "github.com/gofiber/fiber/v3"
)

type ctxKey string

const AuthHeader = "Authorization"
const AuthKey ctxKey = "auth"

type TokenParser interface {
    ParseAccessToken(ctx context.Context, token string) (*domain.Auth, error)
}

func ValidateAuth(tokenParser TokenParser) fiber.Handler {
    return func(ctx fiber.Ctx) error {
        authHeader := ctx.Get(AuthHeader)
        if authHeader == "" {
            log.Printf("missing authorization header")
            
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Missing Authorization Header",
            })
        }

        token := parseBearerToken(authHeader)
        if token == "" {
            log.Printf("invalid authorization format")
            
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Invalid Authorization Format",
            })
        }
        
        auth, err := tokenParser.ParseAccessToken(ctx.Context(), token)
        if err != nil {
            log.Printf("failed to parse token: %v", err)
            
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": "Invalid Access Token",
            })
        }

        ctx.Locals(string(AuthKey), auth)
        return ctx.Next()
    }
}

func GetAuth(ctx fiber.Ctx) (*domain.Auth, bool) {
    val := ctx.Locals(string(AuthKey))
    if val == nil {
        return nil, false
    }
    
    auth, ok := val.(*domain.Auth)
    return auth, ok
}

func parseBearerToken(header string) string {
    parts := strings.SplitN(header, " ", 2)
    if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
        return ""
    }
    return parts[1]
}
