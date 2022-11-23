package exception

import (
	"errors"

	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	code := fiber.StatusInternalServerError

	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if code >= fiber.StatusInternalServerError {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorRes{
			Status:  "failed",
			Message: "internal server error",
		})
	}

	errResult := ctx.Status(code).JSON(response.ErrorRes{
		Status:  "failed",
		Message: err.Error(),
	})

	if errResult != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorRes{
			Status:  "failed",
			Message: "internal server error",
		})
	}

	return nil
}