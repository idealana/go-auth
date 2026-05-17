package handler

import (
	"github.com/gofiber/fiber/v3"
)

type AuthHandlerInterface interface {
	Routes(app *fiber.App)
    Login(ctx fiber.Ctx) error
}
