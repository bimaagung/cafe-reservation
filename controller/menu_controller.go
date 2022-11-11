package controller

import (
	"github.com/bimaagung/cafe-reservation/exception"
	"github.com/bimaagung/cafe-reservation/model"
	"github.com/bimaagung/cafe-reservation/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Menu struct {
	MenuUseCase usecase.MenuUseCase
}

func NewMenuController(menuUseCase *usecase.MenuUseCase) Menu {
	return Menu{MenuUseCase: *menuUseCase}
}

func (controller *Menu) Route(app *fiber.App) {
	app.Get("/api/menu", controller.GetList)
	app.Post("/api/menu", controller.Insert)
	app.Delete("/api/menu/:id", controller.Delete)
}

func (controller *Menu) Insert(c *fiber.Ctx) error {
	request := model.Menu{}
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.Error(err)

	response := controller.MenuUseCase.Add(request)
	return c.JSON(model.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: response,
	})
}

func (controller *Menu) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	response := controller.MenuUseCase.Delete(id)
	return c.JSON(model.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: response,
	})
}

func (controller *Menu) GetList(c *fiber.Ctx) error {
	response := controller.MenuUseCase.GetList()
	return c.JSON(model.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: response,
	})
}