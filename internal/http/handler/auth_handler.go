package handler

import (
	"github.com/gofiber/fiber/v3"
	"go-auth/internal/model/web"
	"go-auth/internal/service"
)

func NewAuthHandler(service *service.AuthService) *AuthHandler {
	return &AuthHandler{service}
}

type AuthHandler struct {
	service *service.AuthService
}

func (handler *AuthHandler) Routes(app *fiber.App) {
	app.Post("/login", handler.Login)
}

func (handler *AuthHandler) Login(c fiber.Ctx) error {
	var req web.LoginRequest

	if err := c.Bind().Body(&req); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid Request.",
        })
    }

    user, err := handler.service.Login(req)

    if err != nil {
    	return c.JSON(fiber.Map{
	        "message": err.Error(),
	    })
    }

    return c.JSON(fiber.Map{
        "message": "Login Successfully.",
        "data": user,
    })
}
