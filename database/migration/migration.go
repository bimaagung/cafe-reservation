package main

import (
	menudomain "github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/bimaagung/cafe-reservation/pkg/dotenv"
	postgresdb "github.com/bimaagung/cafe-reservation/pkg/postgres"
	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
)

func init(){
	dotenv.LoadEnvVariables()
}

func main() {
	database := postgresdb.NewPostgresDB()
	database.AutoMigrate(&menudomain.Menu{}, &userdomain.User{})
}