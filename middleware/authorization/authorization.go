package authorization

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/bimaagung/cafe-reservation/utils/response"
	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt"
)

func AuthValidate(c *gin.Context) {

	var claimsToken jwt.MapClaims
	var getToken string
	authorization := c.Request.Header["Authorization"]

	log.Println(authorization)

	if len(authorization) < 1 {
		c.JSON(401, response.ErrorRes{
			Status: "fail",
			Message: "unauthorized 1",
		})
		return
	}

	splitString := strings.Split(authorization[0], " ")

	if(len(splitString) < 2) {
		getToken = splitString[0]
	}else{
		getToken = splitString[1]
	}

	token, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: error")
		}

		return []byte(os.Getenv("ACCESS_TOKEN_KEY")), nil
	})

	if err != nil {
		c.JSON(401, response.ErrorRes{
			Status: "fail",
			Message: "unauthorized 2",
		})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims);

	if ok && token.Valid {
		claimsToken = claims
	}else {
		c.JSON(401, response.ErrorRes{
			Status: "fail",
			Message: "unauthorized 3",
		})
		return
	}
	
	// pasing value
	c.Set("user", claimsToken)
	
	c.Next()
}