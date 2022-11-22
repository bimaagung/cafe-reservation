package dotenv

import (
	"log"

	"github.com/bimaagung/cafe-reservation/utils/exception"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		exception.CheckError(err)
	}
}