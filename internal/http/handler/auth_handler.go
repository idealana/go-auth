package handler

import (
    "errors"
    "log/slog"

    "go-auth/internal/apperror"
	"go-auth/internal/dto"
    "go-auth/internal/http/middleware"
    "go-auth/internal/service"

    "github.com/gofiber/fiber/v3"
)

func NewAuthHandler(
    authService service.AuthServiceInterface,
    rv *middleware.RequestValidator,
) AuthHandlerInterface {
	return &AuthHandler{authService, rv}
}

type AuthHandler struct {
	authService service.AuthServiceInterface
    reqValidator *middleware.RequestValidator
}

func (handler *AuthHandler) Routes(app *fiber.App) {
    app.Post("/login",
        middleware.ValidateRequest[dto.LoginRequest](handler.reqValidator),
        handler.Login,
    )
}

func (handler *AuthHandler) Login(ctx fiber.Ctx) error {
	req, ok := middleware.GetRequest[dto.LoginRequest](ctx)
    if !ok {
        slog.Error("validated request not found in context",
            "handler", "AuthHandler.Login",
            "path", ctx.Path(),
        )
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Failed to parse request",
        })
    }

    result, err := handler.authService.Login(ctx.Context(), *req)

    if err != nil {
        if errors.Is(err, apperror.ErrInvalidCredentials) {
            return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
                "message": err.Error(),
            })
        }
    	
        return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "message": "Internal Service Error",
        })
    }

    data := dto.LoginResponse{
        UserID: result.UserID,
        AccessToken: result.AccessToken,
    }

    return ctx.JSON(fiber.Map{
        "message": "Login Successfully",
        "data": data,
    })
}
