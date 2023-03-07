package http

import (
	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/middleware/authorization"
	"github.com/gin-gonic/gin"

	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/google/uuid"
)

// Membawa usecase ke controller dan disimpan ke struct Menu
func NewMenuController(menuUseCase *domain.MenuUseCase) Menu {
	return Menu{MenuUseCase: *menuUseCase}
}

// Tempat menyimpan menu usecase
type Menu struct {
	MenuUseCase domain.MenuUseCase
}

func (controller *Menu) Route(app *gin.Engine) {
	app.GET("/api/menu", controller.GetList)
	app.GET("/api/menu/:id", controller.GetById)
	app.POST("/api/menu", authorization.AuthValidate, controller.Insert)
	app.PUT("/api/menu/:id", authorization.AuthValidate, controller.Update)
	app.DELETE("/api/menu/:id", authorization.AuthValidate, controller.Delete)
}


// AddMenu godoc
// @Schemes
// @Security ApiKeyAuth
// @Description add menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param   name  formData     string     true  "name menu"       example(Cappucino)
// @Param   price  formData     int     true  "price menu"       example(15000)
// @Param   stock  formData     int     true  "count stock menu"       example(10)
// @Param image formData file true "Upload cover image menu"
// @Success 200 {object} response.SuccessRes{data=domain.MenuRes}
// @Router /api/menu [post]
func (controller *Menu) Insert(c *gin.Context) {
	var request domain.MenuReq
	request.Id = uuid.New().String()
	if err := c.Bind(&request); err != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: "error response",
		})
		return
	}

	file, errFile := c.FormFile("image")
	if errFile != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: errFile.Error(),
		})
		return
	}
	
	request.File = file
	
	result, err := controller.MenuUseCase.Add(c, &request)

	if err != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: errFile.Error(),
		})
		return
	}

	c.JSON(200, response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

// DeleteMenu godoc
// @Schemes
// @Security ApiKeyAuth
// @Description delete menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param id   path string true "Menu ID"
// @Success 200 {object} response.SuccessRes{status=string,message=string}
// @Router /api/menu/{id} [delete]
func (controller *Menu) Delete(c *gin.Context) {
	id := c.Param("id")

	_, err := controller.MenuUseCase.Delete(c, id)

	if err != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.SuccessRes{
		Status: "ok",
		Message: "success",
	})
}

// GetListMenu godoc
// @Schemes
// @Description get list menu
// @Tags Menu
// @Accept json
// @Produce json
// @Success 200 {object} response.SuccessRes{data=[]domain.MenuRes}
// @Router /api/menu [get]
func (controller *Menu) GetList(c *gin.Context) {

	result, err := controller.MenuUseCase.GetList(c)

	if err != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}

// GetByIdMenu godoc
// @Schemes
// @Description get menu by id
// @Tags Menu
// @Accept json
// @Produce json
// @Param id   path string true "Menu ID"
// @Success 200 {object} response.SuccessRes{data=domain.MenuRes}
// @Router /api/menu/{id} [get]
func (controller *Menu) GetById(c *gin.Context) {
	 id := c.Param("id")

	 result, err := controller.MenuUseCase.GetById(c ,id)

	 if err != nil {
		c.JSON(200, response.ErrorRes{
			Status: "ok",
			Message: err.Error(),
	 	})
	}

	 c.JSON(200, response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	 })
}


// UpdateMenu godoc
// @Schemes
// @Security ApiKeyAuth
// @Description update menu
// @Tags Menu
// @Accept json
// @Produce json
// @Param id   path string true "Menu ID"
// @Param   name  formData     string     true  "name menu"       example(Cappucino)
// @Param   price  formData     int     true  "price menu"       example(15000)
// @Param   stock  formData     int     true  "count stock menu"       example(10)
// @Param image formData file true "Upload cover image menu"
// @Success 200 {object} response.SuccessRes{data=domain.MenuRes}
// @Router /api/menu/{id} [put]
func (controller *Menu) Update(c *gin.Context) {
	var request domain.MenuReq
	id := c.Param("id")

	if err := c.Bind(&request); err != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: "error response",
		})
		return
	}

	file, _ := c.FormFile("image")
	
	request.File = file

	result, err := controller.MenuUseCase.Update(c, id, &request)

	if err != nil {
		c.JSON(400, response.ErrorRes{
			Status: "fail",
			Message: err.Error(),
		})
	}

	c.JSON(200, response.SuccessRes{
		Status: "ok",
		Message: "success",
		Data: result,
	})
}
