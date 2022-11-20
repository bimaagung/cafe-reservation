package controller

import (
	"fmt"

	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/bimaagung/cafe-reservation/menu/usecase"
	"github.com/bimaagung/cafe-reservation/middleware/authorization"

	"github.com/bimaagung/cafe-reservation/utils/exception"
	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Membawa usecase ke controller dan disimpan ke struct Menu
func NewMenuController(menuUseCase *usecase.MenuUseCase) Menu {
	return Menu{MenuUseCase: *menuUseCase}
}

// Tempat menyimpan menu usecase
type Menu struct {
	MenuUseCase usecase.MenuUseCase
}

func (controller *Menu) Route(app *fiber.App) {
	app.Get("/api/menu", controller.GetList)
	app.Get("/api/menu/:id", controller.GetById)
	app.Post("/api/menu", authorization.AuthValidate, controller.Insert)
	app.Put("/api/menu/:id", authorization.AuthValidate, controller.Update)
	app.Delete("/api/menu/:id", authorization.AuthValidate, controller.Delete)
}

func (controller *Menu) Insert(ctx *fiber.Ctx) error {
	var request domain.MenuReq

	request.Id = uuid.New().String()
	if err := ctx.BodyParser(&request); err != nil {
		panic(exception.ClientError{
			Message: err.Error(),
		})
	}

	file, errFile := ctx.FormFile("image")
	if errFile != nil {
		panic(exception.ClientError{
			Message: errFile.Error(),
		})
	}
	
	request.File = file
	
	result := controller.MenuUseCase.Add(ctx, request)

	return ctx.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

func (controller *Menu) Delete(ctx *fiber.Ctx) error {
	
	user := ctx.Locals("user")
	fmt.Println(user)

	id := ctx.Params("id")

	result := controller.MenuUseCase.Delete(ctx, id)
	return ctx.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

func (controller *Menu) GetList(ctx *fiber.Ctx) error {
	result := controller.MenuUseCase.GetList(ctx)
	return ctx.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

func (controller *Menu) GetById(ctx *fiber.Ctx) error {
	 id := ctx.Params("id")

	 result := controller.MenuUseCase.GetById(ctx ,id)

	 return ctx.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	 })
}

func (controller *Menu) Update(ctx *fiber.Ctx) error {
	var request domain.MenuReq
	id := ctx.Params("id")
	if err := ctx.BodyParser(&request); err != nil {
		panic(exception.ClientError{
			Message: err.Error(),
		})
	}

	result := controller.MenuUseCase.Update(ctx, id, request)

	return ctx.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}
