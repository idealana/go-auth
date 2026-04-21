package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v3"
    "go-auth/internal/http/handler"
    "go-auth/internal/service"
    "go-auth/internal/repository"
    "go-auth/pkg/validator"
    "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
    userRepository := repository.NewUserRepository()
    
    authService := service.NewAuthService(userRepository)
    authHandler := handler.NewAuthHandler(authService)

    appConfig := fiber.Config{
    	AppName: os.Getenv("APP_NAME"),
        StructValidator: validator.NewValidator(),
    }

	app := fiber.New(appConfig)

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
