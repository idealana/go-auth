package main

import (
	"log"
	"github.com/gofiber/fiber/v3"
    "go-auth/internal/helper"
    "go-auth/internal/http/handler"
    "go-auth/internal/service"
    "go-auth/internal/repository"
    "go-auth/pkg/validator"
    "go-auth/pkg/utils"
    "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	
	if err != nil {
		log.Fatal("Error loading .env file")
	}

    jwtAccessKey := helper.GetEnvString("JWT_ACCESS_KEY", "!JWTAccessKey!")
    jwtAccessExpired := helper.GetEnvInt("JWT_ACCESS_EXPIRED", 15)
    jwt := utils.NewJWT(jwtAccessKey, jwtAccessExpired)
	
    userRepository := repository.NewUserRepository()
    
    authService := service.NewAuthService(userRepository, jwt)
    authHandler := handler.NewAuthHandler(authService)

    appConfig := fiber.Config{
    	AppName: helper.GetEnvString("APP_NAME", "Go App"),
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
