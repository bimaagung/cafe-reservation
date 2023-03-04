package http

import (
	"fmt"

	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/middleware/authorization"

	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Membawa usecase ke controller dan disimpan ke struct Menu
func NewMenuController(menuUseCase *domain.MenuUseCase) Menu {
	return Menu{MenuUseCase: *menuUseCase}
}

// Tempat menyimpan menu usecase
type Menu struct {
	MenuUseCase domain.MenuUseCase
}

func (controller *Menu) Route(app *fiber.App) {
	app.Get("/api/menu", controller.GetList)
	app.Get("/api/menu/:id", controller.GetById)
	app.Post("/api/menu", authorization.AuthValidate, controller.Insert)
	app.Put("/api/menu/:id", authorization.AuthValidate, controller.Update)
	app.Delete("/api/menu/:id", authorization.AuthValidate, controller.Delete)
}

func (controller *Menu) Insert(c *fiber.Ctx) error {
	var ctx = c.Context()

	var request domain.MenuReq
	request.Id = uuid.New().String()
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, errFile.Error())
	}
	
	request.File = file
	
	result, err := controller.MenuUseCase.Add(ctx, &request)

	if err != nil {
		return err
	}

	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

func (controller *Menu) Delete(c *fiber.Ctx) error {
	var ctx = c.Context()

	user := c.Locals("user")
	fmt.Println(user)

	id := c.Params("id")

	_, err := controller.MenuUseCase.Delete(ctx, id)

	if err != nil {
		return err
	}

	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
	})
}

func (controller *Menu) GetList(c *fiber.Ctx) error {
	var ctx = c.Context()

	result, err := controller.MenuUseCase.GetList(ctx)

	if err != nil {
		return err
	}

	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

func (controller *Menu) GetById(c *fiber.Ctx) error {
	var ctx = c.Context()

	 id := c.Params("id")

	 result, err := controller.MenuUseCase.GetById(ctx ,id)

	 if err != nil {
		return err
	}

	 return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	 })
}

func (controller *Menu) Update(c *fiber.Ctx) error {
	var ctx = c.Context()

	var request domain.MenuReq
	id := c.Params("id")
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
	}

	file, _ := c.FormFile("image")
	
	request.File = file

	result, err := controller.MenuUseCase.Update(ctx, id, &request)

	if err != nil {
		return err
	}

	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}
