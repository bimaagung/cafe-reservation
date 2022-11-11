package http

import (
	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/exception"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Menu struct {
	MenuUseCase domain.MenuUseCase
}

func NewMenuController(menuUseCase *domain.MenuUseCase) Menu {
	return Menu{MenuUseCase: *menuUseCase}
}

func (controller *Menu) Route(app *fiber.App) {
	app.Get("/api/menu", controller.GetList)
	app.Post("/api/menu", controller.Insert)
	app.Delete("/api/menu/:id", controller.Delete)
}

func (controller *Menu) Insert(c *fiber.Ctx) error {
	request := domain.Menu{}
	err := c.BodyParser(&request)
	request.Id = uuid.New().String()

	exception.Error(err)

	response := controller.MenuUseCase.Add(request)
	return c.JSON(domain.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: response,
	})
}

func (controller *Menu) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	response := controller.MenuUseCase.Delete(id)
	return c.JSON(domain.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: response,
	})
}

func (controller *Menu) GetList(c *fiber.Ctx) error {
	response := controller.MenuUseCase.GetList()
	return c.JSON(domain.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: response,
	})
}