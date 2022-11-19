package exception

import (
	"log"

	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, notFound := err.(NotFoundError)
	_, clientError := err.(ClientError)
	_, unauthorized := err.(Unathorized)
	

	if clientError {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.ErrorRes{
			Status: "failed",
			Message: err.Error(),
		})
	}

	if notFound {
		return ctx.Status(fiber.StatusNotFound).JSON(response.ErrorRes{
			Status: "failed",
			Message: err.Error(),
		})
	}

	if unauthorized {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.ErrorRes{
			Status: "failed",
			Message: "unauthorized",
		})
	}

	log.Fatal(err.Error())
	return ctx.JSON(response.ErrorRes{
		Status: "failed",
		Message: "Internal Server Error",
	})
}