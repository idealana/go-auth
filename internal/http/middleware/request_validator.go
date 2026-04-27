package middleware

import (
	"errors"
	"log"

	"go-auth/pkg/validator"

	"github.com/gofiber/fiber/v3"
)

const RequestKey = "validated_request"

func ValidateRequest[T any]() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var req T
		
		if err := ctx.Bind().Body(&req); err != nil {
			var valErr validator.ValidationError
			if errors.As(err, &valErr) {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"message": "Bad Request",
					"errors": valErr.Errors,
				})
			}
			
			log.Printf("invalid request: %v", err)
			
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Invalid Request",
				"errors": map[string]string{},
			})
		}
		
		ctx.Locals(RequestKey, &req)
		return ctx.Next()
	}
}

func GetRequest[T any](ctx fiber.Ctx) (*T, bool) {
	val := ctx.Locals(RequestKey)
	if val == nil {
		return nil, false
	}
	
	req, ok := val.(*T)
	return req, ok
}
