package postgresdb

import (
	"os"

	"github.com/bimaagung/cafe-reservation/utils/exception"

	postgres "go.elastic.co/apm/module/apmgormv2/v2/driver/postgres"
	"gorm.io/gorm"
)


func NewPostgresDB() *gorm.DB {
	dsn := os.Getenv("DB_URL")
	DB, err :=  gorm.Open(postgres.Open(dsn), &gorm.Config{})
	exception.Error(err)

	return DB
}


