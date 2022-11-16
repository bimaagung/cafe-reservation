package main

import (
	"github.com/bimaagung/cafe-reservation/config"
	"github.com/bimaagung/cafe-reservation/models/domain"
)

func init(){
	config.LoadEnvVariables()
}

func main() {
	database := config.NewPostgresDB()
	database.AutoMigrate(&domain.Menu{})
}