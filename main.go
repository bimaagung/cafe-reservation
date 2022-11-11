package main

import (
	"github.com/bimaagung/cafe-reservation/config"
	"github.com/bimaagung/cafe-reservation/controller"
	"github.com/bimaagung/cafe-reservation/exception"
	"github.com/bimaagung/cafe-reservation/repository"
	"github.com/bimaagung/cafe-reservation/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init(){
	config.LoadEnvVariables()
}

func main() {
	dbPostgres := config.NewPostgresDB()

	menuRepository := repository.NewConnectDB(dbPostgres)
	menuUseCase := usecase.NewMenuUC(&menuRepository)
	menuController := controller.NewMenuController(&menuUseCase)

	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ErrorHandler,
		},
	)
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	menuController.Route(app)

	err := app.Listen(":3000")
	exception.Error(err)
}