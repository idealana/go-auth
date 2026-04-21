package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v3"
    "go-auth/internal/http/handler"
    "go-auth/internal/service"
    "go-auth/internal/repository"
    "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appName := os.Getenv("APP_NAME")
	
    userRepository := repository.NewUserRepository()
    
    authService := service.NewAuthService(userRepository)
    authHandler := handler.NewAuthHandler(authService)

	app := fiber.New(fiber.Config{
		AppName: appName,
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
