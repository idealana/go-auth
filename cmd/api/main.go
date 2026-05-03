package main

import (
    "fmt"
	"log"
    "time"

    "go-auth/internal/config"
    "go-auth/internal/http/handler"
    "go-auth/internal/repository"
    "go-auth/internal/security"
    "go-auth/internal/service"
    "go-auth/pkg/validator"

    "github.com/gofiber/fiber/v3"
    "github.com/joho/godotenv"
)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Println("no .env file found, using system environment variables")
    }

    appConfig := fiber.Config{
        AppName: config.GetAppName(),
        StructValidator: validator.NewValidator(),
    }

    jwtAccessKey, err := config.GetJWTAccessKey()
    if err != nil {
        log.Fatalf("config error: %v", err)
    }

    jwtAccessExpired := time.Minute * time.Duration(config.GetJWTAccessExpired())
    jwtAuth, err := security.NewJWTAuthToken(appConfig.AppName, jwtAccessKey, jwtAccessExpired)
    if err != nil {
        log.Fatalf("failed to initialize jwt: %v", err)
    }

    bcryptPassword := security.NewBcryptPassword()
	
    // REPOSITORIES
    userRepository := repository.NewUserRepository()
    
    // SERVICES
    authService := service.NewAuthService(userRepository, jwtAuth, bcryptPassword)

    // HANDLERS
    authHandler := handler.NewAuthHandler(authService)

    app := fiber.New(appConfig)

    // ROUTES
	app.Get("/", func(c fiber.Ctx) error {
        return c.JSON(fiber.Map{
        	"message": "Hello World!",
        })
    })

    authHandler.Routes(app)
    // END ROUTES

    port := config.GetAppPort()

    if err := app.Listen(fmt.Sprintf(":%d", port)); err != nil {
        log.Fatalf("server failed to start: %v", err)
    }
}
