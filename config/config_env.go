package config

import (
	"log"

	"github.com/bimaagung/cafe-reservation/exception"
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		exception.Error(err)
	}
}