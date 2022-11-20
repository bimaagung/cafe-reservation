package main

import (
	"github.com/bimaagung/cafe-reservation/pkg/dotenv"
	postgresdb "github.com/bimaagung/cafe-reservation/pkg/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	// menu
	menucontroller "github.com/bimaagung/cafe-reservation/menu/controller"
	menurepository "github.com/bimaagung/cafe-reservation/menu/repository"
	menuusecase "github.com/bimaagung/cafe-reservation/menu/usecase"
	"github.com/bimaagung/cafe-reservation/utils/exception"

	// user
	usercontroller "github.com/bimaagung/cafe-reservation/user/controller"
	userrepository "github.com/bimaagung/cafe-reservation/user/repository"
	userusecase "github.com/bimaagung/cafe-reservation/user/usecase"

	"go.elastic.co/apm/module/apmfiber/v2"
)

func init(){
	dotenv.LoadEnvVariables()
}

func main() {
	dbPostgres := postgresdb.NewPostgresDB()

	menuRepository := menurepository.NewConnectDB(dbPostgres)
	menuUseCase := menuusecase.NewMenuUC(&menuRepository)
	menuController := menucontroller.NewMenuController(&menuUseCase)

	userRepository := userrepository.NewUserRepository(dbPostgres)
	userUseCase := userusecase.NewUserUC(&userRepository)
	userController := usercontroller.NewUserController(&userUseCase)
	

	app := fiber.New(
		fiber.Config{
			ErrorHandler: exception.ErrorHandler,
		},
	)

	app.Use(apmfiber.Middleware())
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	menuController.Route(app)
	userController.Route(app)

	err := app.Listen(":3000")
	exception.Error(err)
}