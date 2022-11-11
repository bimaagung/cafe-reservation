package main

import (
	"github.com/bimaagung/cafe-reservation/config"
	"github.com/bimaagung/cafe-reservation/entity"
)

func init(){
	config.LoadEnvVariables()
}

func main() {
	database := config.NewPostgresDB()
	database.AutoMigrate(&entity.Menu{})
}