package exception

import (
	"fmt"

	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	errNotFound, notFound := err.(NewNotFoundError)
	errClientError, clientError := err.(NewClientError)
	_, unauthorized := err.(NewUnauthorized)

	if clientError {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorRes{
			Status: "failed",
			Message: errClientError.Message,
		})
	} else if notFound {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorRes{
			Status: "failed",
			Message: errNotFound.Message,
		})
	} else if unauthorized {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorRes{
			Status: "failed",
			Message: "unauthorized",
		})
	}else{
		fmt.Errorf("internal error:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.ErrorRes{
			Status: "failed",
			Message: "internal server error",
		})
	}
}