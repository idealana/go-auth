package main

import (
	"github.com/gofiber/fiber/v3"
    "go-auth/internal/http/handler"
    "go-auth/internal/service"
    "go-auth/internal/repository"
    "go-auth/pkg/validator"
)

func main() {
    userRepository := repository.NewUserRepository()
    
    authService := service.NewAuthService(userRepository)
    authHandler := handler.NewAuthHandler(authService)

    customValidator := validator.NewValidator()

	app := fiber.New(fiber.Config{
        StructValidator: customValidator,
    })

    // ROUTES
	app.Get("/", func(c fiber.Ctx) error {
        return c.JSON(fiber.Map{
        	"message": "Hello World!",
        })
    })

    authHandler.Routes(app)
    // END ROUTES

    app.Listen(":8100")
}
