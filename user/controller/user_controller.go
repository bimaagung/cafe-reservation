package controller

import (
	"github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/bimaagung/cafe-reservation/user/usecase"
	"github.com/bimaagung/cafe-reservation/utils/exception"
	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func NewUserController(userUseCae *usecase.UserUseCase) User {
	return User{UserUseCase: *userUseCae}
}

type User struct {
	UserUseCase usecase.UserUseCase
}

func (controller *User) Route(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/auth/register", controller.Register)
}

func (controller *User) Register(c *fiber.Ctx) error {
	
	var request domain.UserReq

	request.Id = uuid.New().String()
	
	if err := c.BodyParser(&request); err != nil {
		panic(exception.ClientError{
			Message: err.Error(),
		})
	}

	result := controller.UserUseCase.Create(request)

	

	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})	
}