package tokenmanager

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(claims jwt.MapClaims) (string, error) {
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_KEY")))
	
	if err != nil {
		return "", err
	}
	
	return t, nil
}