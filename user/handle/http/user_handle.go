package http

import (
	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewUserController(userUseCae domain.UserUseCase) User {
	return User{UserUseCase: userUseCae}
}

type User struct {
	UserUseCase domain.UserUseCase
}

func (controller *User) Route(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/register", controller.Register)
}

func (controller *User) Register(c *fiber.Ctx) error {
	
	var ctx = c.Context()
	var request domain.UserReq

	request.Id = uuid.New().String()
	
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	result, err := controller.UserUseCase.Create(ctx, &request)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})	
}