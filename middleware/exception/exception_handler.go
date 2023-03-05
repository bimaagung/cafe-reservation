package exception

import (
	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gin-gonic/gin"
)


func ErrorHandler(c *gin.Context, err any){
	c.AbortWithStatusJSON(500, response.ErrorRes{
		Status: "fail",
		Message: "internal server error",
	})
}