package controller

import (
	"os"
	"strconv"
	"time"

	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/bimaagung/cafe-reservation/menu/usecase"
	minioUpload "github.com/bimaagung/cafe-reservation/pkg/minio"
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
	app.Post("/api/menu", controller.Insert)
	app.Put("/api/menu/:id", controller.Update)
	app.Delete("/api/menu/:id", controller.Delete)
}

func (controller *Menu) Insert(c *fiber.Ctx) error {
	var request domain.MenuReq

	request.Id = uuid.New().String()
	if err := c.BodyParser(&request); err != nil {
		panic(exception.ClientError{
			Message: err.Error(),
		})
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		panic(exception.ClientError{
			Message: errFile.Error(),
		})
	}

	
	bucketName := "menu"
	timestamp := time.Now().Unix()
	objectName :=  strconv.FormatInt(timestamp, 16) +"-"+ file.Filename
	
	// Upload file menggunakan Minio
	errUpload := minioUpload.UploadFile(file, bucketName, objectName)

	if errUpload != nil {
		panic(exception.ClientError{
			Message: errFile.Error(),
		})
	}

	urlFile := os.Getenv("MINIO_URL_FILE")+"/"+bucketName+"/"+objectName
	
	result := controller.MenuUseCase.Add(request, urlFile)

	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

func (controller *Menu) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	result := controller.MenuUseCase.Delete(id)
	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

func (controller *Menu) GetList(c *fiber.Ctx) error {
	result := controller.MenuUseCase.GetList()
	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

func (controller *Menu) GetById(c *fiber.Ctx) error {
	 id := c.Params("id")

	 result := controller.MenuUseCase.GetById(id)

	 return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	 })
}

func (controller *Menu) Update(c *fiber.Ctx) error {
	var request domain.MenuReq
	id := c.Params("id")
	if err := c.BodyParser(&request); err != nil {
		panic(exception.ClientError{
			Message: err.Error(),
		})
	}

	result := controller.MenuUseCase.Update(id, request)

	return c.JSON(response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}
