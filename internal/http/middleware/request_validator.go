package middleware

import (
	"errors"
	"fmt"

	"go-auth/internal/http/response"
	"go-auth/internal/logger"
	"go-auth/pkg/validator"

	"github.com/gofiber/fiber/v3"
)

type requestKey struct {
	typeName string
}

type RequestValidator struct {
	log logger.Logger
}

func NewRequestValidator(log logger.Logger) *RequestValidator {
	return &RequestValidator{
		log: log,
	}
}

func ValidateRequest[T any](rv *RequestValidator) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var req T

		if err := ctx.Bind().Body(&req); err != nil {
			return rv.handleBindError(ctx, err)
		}
		
		ctx.Locals(requestKey{typeName: typeNameOf[T]()}, &req)
		return ctx.Next()
	}
}

func GetRequest[T any](ctx fiber.Ctx) (*T, bool) {
	val := ctx.Locals(requestKey{typeName: typeNameOf[T]()})
	if val == nil {
		return nil, false
	}
	
	req, ok := val.(*T)
	return req, ok
}

func (r *RequestValidator) handleBindError(ctx fiber.Ctx, err error) error {
	var valErr *validator.ValidationError
	if errors.As(err, &valErr) {
		return response.BadRequest(ctx, "Bad Request", valErr.Errors)
	}

	r.log.Error("failed to bind request body", "error", err)
	return response.BadRequest(ctx, "Invalid Request", map[string]string{})
}

func typeNameOf[T any]() string {
	return fmt.Sprintf("%T", *new(T))
}
