package main

import (
	"github.com/bimaagung/cafe-reservation/config"
	"github.com/bimaagung/cafe-reservation/exception"
	"github.com/bimaagung/cafe-reservation/menu/delivery/http"
	"github.com/bimaagung/cafe-reservation/menu/repository/postgres"
	"github.com/bimaagung/cafe-reservation/menu/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init(){
	config.LoadEnvVariables()
}

func main() {
	dbPostgres := config.NewPostgresDB()

	menuRepository := postgres.NewConnectDB(dbPostgres)
	menuUseCase := usecase.NewMenuUC(&menuRepository)
	menuController := http.NewMenuController(&menuUseCase)

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