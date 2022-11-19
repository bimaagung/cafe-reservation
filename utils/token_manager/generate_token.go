package tokenmanager

import (
	"os"

	"github.com/bimaagung/cafe-reservation/utils/exception"
	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(claims jwt.MapClaims) (t string) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))

	if err != nil {
		exception.Error(err.Error())
	}

	return t
}