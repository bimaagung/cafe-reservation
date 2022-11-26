package authorization

import (
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt"
)

func AuthValidate(c *fiber.Ctx) error {

	var claimsToken jwt.MapClaims
	var getToken string
	authorization := c.Get("Authorization")

	if authorization == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	splitString := strings.Split(authorization, " ")

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
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims);

	if ok && token.Valid {
		claimsToken = claims
	}else {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	c.Locals("user", claimsToken)
	
	return c.Next()
}