package main

import (
    "fmt"
    "time"

    "go-auth/internal/config"
    "go-auth/internal/http/handler"
    "go-auth/internal/http/middleware"
    "go-auth/internal/logger"
    "go-auth/internal/repository"
    "go-auth/internal/security"
    "go-auth/internal/service"
    "go-auth/pkg/validator"

    "github.com/gofiber/fiber/v3"
    "github.com/joho/godotenv"
)

func main() {
    log := logger.Resolve()

    if err := godotenv.Load(); err != nil {
        log.Info("no .env file found, using system environment variables")
    }

    validatorService, err := validator.NewValidator()
    if err != nil {
        log.Fatal("failed to initialize validator", "error", err)
    }

    appConfig := fiber.Config{
        AppName: config.GetAppName(),
        StructValidator: validatorService,
    }

    jwtAccessKey, err := config.GetJWTAccessKey()
    if err != nil {
        log.Fatal("jwt config error", "error", err)
    }

    jwtAccessExpired := time.Minute * time.Duration(config.GetJWTAccessExpired())
    jwtAuth, err := security.NewJWTAuthToken(appConfig.AppName, jwtAccessKey, jwtAccessExpired)
    if err != nil {
        log.Fatal("failed to initialize jwt", "error", err)
    }

    reqValidator := middleware.NewRequestValidator(log)
    bcryptPassword := security.NewBcryptPassword()
	
    // REPOSITORIES
    userRepository := repository.NewUserRepository()
    
    // SERVICES
    authService := service.NewAuthService(userRepository, jwtAuth, bcryptPassword)

    // HANDLERS
    authHandler := handler.NewAuthHandler(
        authService,
        reqValidator,
    )

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
        log.Fatal("server failed to start", "error", err)
    }
}
