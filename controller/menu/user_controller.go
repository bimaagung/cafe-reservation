package controller

import (
	"github.com/bimaagung/cafe-reservation/exception"
	"github.com/bimaagung/cafe-reservation/models/web"
	usecase "github.com/bimaagung/cafe-reservation/usecase/user"
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
	
	var request web.UserReq

	request.Id = uuid.New().String()
	
	if err := c.BodyParser(&request); err != nil {
		panic(exception.ClientError{
			Message: err.Error(),
		})
	}

	response := controller.UserUseCase.Create(request)

	return c.JSON(web.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: response,
	})	
}