package main

import (
	"github.com/bimaagung/cafe-reservation/domain"
	"github.com/bimaagung/cafe-reservation/pkg/dotenv"
	postgresdb "github.com/bimaagung/cafe-reservation/pkg/postgres"
)

func init(){
	dotenv.LoadEnvVariables()
}

func main() {
	database := postgresdb.NewPostgresDB()
	database.AutoMigrate(&domain.Menu{}, &domain.User{})
}