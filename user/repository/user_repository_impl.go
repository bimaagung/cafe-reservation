package repository

import (
	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/bimaagung/cafe-reservation/utils/exception"
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

func (repository *postgresUserRepository) Create(user userdomain.User) {
	result := repository.DB.Create(&user)
	exception.Error(result.Error)
	
}

func (repository *postgresUserRepository) GetById(id string)(user userdomain.User) {
	repository.DB.First(&user, "id = ?", id)
	return user
}