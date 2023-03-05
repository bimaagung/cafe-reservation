package http

import (
	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewUserController(userUseCae domain.UserUseCase) User {
	return User{UserUseCase: userUseCae}
}

type User struct {
	UserUseCase domain.UserUseCase
}

func (controller *User) Route(app *gin.Engine) {
	api := app.Group("/api")

	api.POST("/auth/register", controller.Register)
}

func (controller *User) Register(c *gin.Context) {
	var request domain.UserReq

	request.Id = uuid.New().String()
	
	if err := c.BindJSON(&request); err != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: err.Error(),
		})
		return
	}

	result, err := controller.UserUseCase.Create(c, &request)

	if err != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: err.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})	
}