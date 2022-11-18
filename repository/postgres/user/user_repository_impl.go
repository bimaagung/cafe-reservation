package userrepository

import (
	"github.com/bimaagung/cafe-reservation/exception"
	"github.com/bimaagung/cafe-reservation/models/domain"
	"gorm.io/gorm"
)


func NewUserRepository(database *gorm.DB) UserRepository {
	return &postgresUserRepository {
		DB: database,
	}
}

type postgresUserRepository struct {
	DB *gorm.DB
}

func (repository *postgresUserRepository) Create(user domain.User) {
	result := repository.DB.Create(&user)
	exception.Error(result.Error)
	
}

func (repository *postgresUserRepository) GetById(id string)(user domain.User) {
	repository.DB.First(&user, "id = ?", id)
	return user
}