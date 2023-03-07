package main

import (
	"log"

	"github.com/bimaagung/cafe-reservation/middleware/exception"
	"github.com/bimaagung/cafe-reservation/pkg/dotenv"
	miniodb "github.com/bimaagung/cafe-reservation/pkg/minio"
	postgresdb "github.com/bimaagung/cafe-reservation/pkg/postgres"
	redisdb "github.com/bimaagung/cafe-reservation/pkg/redis"
	"github.com/gin-gonic/gin"

	// menu
	menuhandle "github.com/bimaagung/cafe-reservation/menu/handle/http"
	menurepositoryminio "github.com/bimaagung/cafe-reservation/menu/repository/minio"
	menurepositorypostgres "github.com/bimaagung/cafe-reservation/menu/repository/postgres"
	menurepositoryredis "github.com/bimaagung/cafe-reservation/menu/repository/redis"
	menuusecase "github.com/bimaagung/cafe-reservation/menu/usecase"

	// user
	userhandle "github.com/bimaagung/cafe-reservation/user/handle/http"
	userrepository "github.com/bimaagung/cafe-reservation/user/repository/postgres"
	userusecase "github.com/bimaagung/cafe-reservation/user/usecase"

	docs "github.com/bimaagung/cafe-reservation/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init(){
	dotenv.LoadEnvVariables()
}

// @title           Cafe Reservation API
// @version         1.0
// @description     This is a cafe reservation server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:3000
// @BasePath  /api/

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	dbPostgres := postgresdb.NewPostgresDB()
	dbRedis := redisdb.NewRedisDB()
	dbMinio := miniodb.MinioConnection()


	// Menu
	menuRepositoryPostgres := menurepositorypostgres.NewConnectDB(dbPostgres)
	menuRepositoryRedis := menurepositoryredis.NewRepositoryRedis(dbRedis)
	menuRepositoryMinio := menurepositoryminio.NewMinioRepository(dbMinio)

	menuUseCase := menuusecase.NewMenuUC(menuRepositoryPostgres, menuRepositoryRedis, menuRepositoryMinio)
	menuController := menuhandle.NewMenuController(&menuUseCase)

	// User
	userRepository := userrepository.NewUserRepository(dbPostgres)
	userUseCase := userusecase.NewUserUC(userRepository)
	userController := userhandle.NewUserController(userUseCase)
	

	app := gin.Default()
	docs.SwaggerInfo.BasePath = ""
	app.Use(gin.CustomRecovery(exception.ErrorHandler))

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler)) // http://localhost:3000/swagger/index.html

	menuController.Route(app)
	userController.Route(app)
	
	err := app.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
	
}