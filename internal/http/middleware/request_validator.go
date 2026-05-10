package middleware

import (
	"errors"
	"log/slog"

	"go-auth/internal/http/response"
	"go-auth/pkg/validator"

	"github.com/gofiber/fiber/v3"
)

type requestKey struct {
	name string
}

var defaultRequestKey = requestKey{"validated_request"}

type Logger interface {
	Error(message string, args ...any)
}

type defaultLogger struct {}

func (logger *defaultLogger) Error(message string, args ...any) {
	slog.Error(message, args...)
}

func ValidateRequest[T any](logger ...Logger) fiber.Handler {
	log := resolveLogger(logger...)

	return func(ctx fiber.Ctx) error {
		var req T

		if err := ctx.Bind().Body(&req); err != nil {
			return handleBindError(ctx, log, err)
		}
		
		ctx.Locals(defaultRequestKey, &req)
		return ctx.Next()
	}
}

func GetRequest[T any](ctx fiber.Ctx) (*T, bool) {
	val := ctx.Locals(defaultRequestKey)
	if val == nil {
		return nil, false
	}
	
	req, ok := val.(*T)
	return req, ok
}

func handleBindError(ctx fiber.Ctx, log Logger, err error) error {
	var valErr *validator.ValidationError
	if errors.As(err, &valErr) {
		return response.BadRequest(ctx, "Bad Request", valErr.Errors)
	}

	log.Error("failed to bind request body", "error", err)
	return response.BadRequest(ctx, "Invalid Request", map[string]string{})
}

func resolveLogger(loggers ...Logger) Logger {
	if len(loggers) > 0 && loggers[0] != nil {
		return loggers[0]
	}
	return &defaultLogger{}
}
