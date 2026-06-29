package handler

import (
    "errors"

    "go-auth/internal/apperror"
    "go-auth/internal/http/middleware"
    "go-auth/internal/http/response"
    "go-auth/internal/model/dto"
    "go-auth/internal/logger"
    "go-auth/internal/service"

    "github.com/gofiber/fiber/v3"
)

func NewAuthHandler(
    authService service.AuthServiceInterface,
    rv *middleware.RequestValidator,
    log logger.Logger,
    authChecker *middleware.Auth,
) AuthHandlerInterface {
	return &AuthHandler{authService, rv, log, authChecker}
}

type AuthHandler struct {
	authService service.AuthServiceInterface
    reqValidator *middleware.RequestValidator
    log logger.Logger
    authChecker *middleware.Auth
}

func (handler *AuthHandler) Routes(app fiber.Router) {
    app.Post("/login",
        middleware.ValidateRequest[dto.LoginRequest](handler.reqValidator),
        handler.Login,
    )

    auth := app.Group("/auth",
        handler.authChecker.ValidateAuth(),
    )

    auth.Get("/profile",
        handler.Profile,
    )
}

func (handler *AuthHandler) Login(ctx fiber.Ctx) error {
	req, ok := middleware.GetRequest[dto.LoginRequest](ctx)
    if !ok {
        handler.log.Error(
            "validated request not found in context",
            "handler", "AuthHandler.Login",
            "path", ctx.Path(),
        )
        return response.InternalServerError(ctx, "Failed to parse request")
    }

    reqInfo := dto.RequestInfo{
        IPAddress: ctx.IP(),
        UserAgent: ctx.Get(fiber.HeaderUserAgent),
    }

    result, err := handler.authService.Login(ctx.Context(), *req, &reqInfo)

    if err != nil {
        if errors.Is(err, apperror.ErrInvalidCredentials) {
            return response.Unauthorized(ctx, err.Error())
        }
    	
        return response.InternalServerError(ctx, "Internal Service Error")
    }

    return response.Success[dto.LoginResponse](
        ctx,
        "Login Successfully",
        dto.LoginResponse{
            UserID: result.UserID,
            AccessToken: result.AccessToken,
            RefreshToken: result.RefreshToken,
        },
    )
}

func (handler *AuthHandler) Profile(ctx fiber.Ctx) error {
    auth, ok := middleware.GetAuth(ctx)
    if !ok {
        return response.Unauthorized(ctx, "Invalid User")
    }
    
    result, err := handler.authService.GetProfile(ctx.Context(), auth.UserID)
    if err != nil {
        if errors.Is(err, apperror.ErrNotFound) {
            return response.NotFound(ctx, err.Error())
        }
    	
        return response.InternalServerError(ctx, "Internal Service Error")
    }

    return response.Success[dto.ProfileResponse](
        ctx,
        "Success",
        dto.ProfileResponse{
            UserID: result.UserID,
            Email: result.Email,
            Role: result.Role,
            Status: result.Status,
        },
    )
}
