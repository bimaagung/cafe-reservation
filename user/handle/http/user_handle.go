package http

import (
	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gin-gonic/gin"
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

// Register godoc
// @Schemes
// @Description register a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.UserReq true "the input account user"
// @Success 200 {object} response.SuccessRes{data=domain.UserRes}
// @Router /api/auth/register [post]
func (controller *User) Register(c *gin.Context) {
	var request domain.UserReq
	
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