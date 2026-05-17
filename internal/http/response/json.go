package response

import (
	"github.com/gofiber/fiber/v3"
)

func Success[T any](ctx fiber.Ctx, message string, data T) error {
    return ctx.Status(fiber.StatusOK).JSON(newSuccessResponse[T](message, data))
}

func Created[T any](ctx fiber.Ctx, message string, data T) error {
    return ctx.Status(fiber.StatusCreated).JSON(newSuccessResponse[T](message, data))
}

func BadRequest(ctx fiber.Ctx, message string, errs map[string]string) error {
    return ctx.Status(fiber.StatusBadRequest).JSON(newErrorResponse(message, errs))
}

func Unauthorized(ctx fiber.Ctx, message string) error {
    return ctx.Status(fiber.StatusUnauthorized).JSON(newErrorResponse(message, nil))
}

func NotFound(ctx fiber.Ctx, message string) error {
    return ctx.Status(fiber.StatusNotFound).JSON(newErrorResponse(message, nil))
}

func InternalServerError(ctx fiber.Ctx, message string) error {
    if message == "" {
        message = "Internal Server Error"
    }
    return ctx.Status(fiber.StatusInternalServerError).JSON(newErrorResponse(message, nil))
}

func normalizeErrors(errs map[string]string) map[string]string {
    if errs != nil {
        return errs
    }
    return map[string]string{}
}

func newSuccessResponse[T any](message string, data T) SuccessResponse[T] {
    return SuccessResponse[T]{
        Success: true,
        Message: message,
        Data: data,
    }
}

func newErrorResponse(message string, errs map[string]string) ErrorResponse {
    return ErrorResponse{
        Success: false,
        Message: message,
        Errors: normalizeErrors(errs),
    }
}
