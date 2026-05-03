package handler

import (
    "errors"

    "go-auth/internal/apperror"
	"go-auth/internal/dto"
    "go-auth/internal/service"
	"go-auth/pkg/validator"

    "github.com/gofiber/fiber/v3"
)

func NewAuthHandler(authService service.AuthServiceInterface) AuthHandlerInterface {
	return &AuthHandler{authService}
}

type AuthHandler struct {
	authService service.AuthServiceInterface
}

func (handler *AuthHandler) Routes(app *fiber.App) {
	app.Post("/login", handler.Login)
}

func (handler *AuthHandler) Login(ctx fiber.Ctx) error {
	var req dto.LoginRequest

	if err := ctx.Bind().Body(&req); err != nil {
        var valErr *validator.ValidationError
		if errors.As(err, &valErr) {
            return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "message": "Bad request",
                "errors": valErr.Errors,
            })
        }

        return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "message": "Invalid Request",
            "errors": map[string]string{},
        })
    }

    result, err := handler.authService.Login(ctx.Context(), req)

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
