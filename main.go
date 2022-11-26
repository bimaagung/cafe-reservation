package main

import (
	"log"

	"github.com/bimaagung/cafe-reservation/pkg/dotenv"
	postgresdb "github.com/bimaagung/cafe-reservation/pkg/postgres"
	redisdb "github.com/bimaagung/cafe-reservation/pkg/redis"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"

	// menu
	menucontroller "github.com/bimaagung/cafe-reservation/menu/controller"
	menurepositorypostgres "github.com/bimaagung/cafe-reservation/menu/repository/postgres"
	menurepositoryredis "github.com/bimaagung/cafe-reservation/menu/repository/redis"
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
	dbRedis := redisdb.NewRedisDB()

	// Menu
	menuRepositoryPostgres := menurepositorypostgres.NewConnectDB(dbPostgres)
	menuRepositoryRedis := menurepositoryredis.NewRepositoryRedis(dbRedis)
	menuUseCase := menuusecase.NewMenuUC(menuRepositoryPostgres, menuRepositoryRedis)
	menuController := menucontroller.NewMenuController(&menuUseCase)

	// User
	userRepository := userrepository.NewUserRepository(dbPostgres)
	userUseCase := userusecase.NewUserUC(userRepository)
	userController := usercontroller.NewUserController(userUseCase)
	

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

	if err := app.Listen(":3000"); err != nil {
		log.Fatal(err)
	}
	
}