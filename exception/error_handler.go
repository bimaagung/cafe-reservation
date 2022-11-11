package exception

import (
	"log"

	"github.com/bimaagung/cafe-reservation/model"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	_, notFound := err.(NotFoundError)
	_, clientError := err.(ClientError)

	

	if clientError {
		return ctx.Status(400).JSON(model.ErrorRes{
			Status: "failed",
			Message: err.Error(),
		})
	}

	if notFound {
		return ctx.Status(404).JSON(model.ErrorRes{
			Status: "failed",
			Message: err.Error(),
		})
	}

	log.Fatal(err.Error())
	return ctx.JSON(model.ErrorRes{
		Status: "failed",
		Message: "Internal Server Error",
	})
}