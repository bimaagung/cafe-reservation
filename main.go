package main

import (
	"github.com/bimaagung/cafe-reservation/config"
	controller "github.com/bimaagung/cafe-reservation/controller/menu"
	"github.com/bimaagung/cafe-reservation/exception"
	menurepository "github.com/bimaagung/cafe-reservation/repository/postgres/menu"
	userrepository "github.com/bimaagung/cafe-reservation/repository/postgres/user"
	menuusecase "github.com/bimaagung/cafe-reservation/usecase/menu"
	userusecase "github.com/bimaagung/cafe-reservation/usecase/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init(){
	config.LoadEnvVariables()
}

func main() {
	dbPostgres := config.NewPostgresDB()

	menuRepository := menurepository.NewConnectDB(dbPostgres)
	menuUseCase := menuusecase.NewMenuUC(&menuRepository)
	menuController := controller.NewMenuController(&menuUseCase)

	userRepository := userrepository.NewUserRepository(dbPostgres)
	userUseCase := userusecase.NewUserUC(&userRepository)
	userController := controller.NewUserController(&userUseCase)
	

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
	userController.Route(app)

	err := app.Listen(":3000")
	exception.Error(err)
}