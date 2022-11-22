package repository

import (
	"context"
	"errors"

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

func (repository *postgresUserRepository) Create(ctx context.Context, user userdomain.User) {
	db := repository.DB.WithContext(ctx)
	result := db.Create(&user)
	exception.CheckError(result.Error)
}

func (repository *postgresUserRepository) GetByUsername(ctx context.Context, username string)(userdomain.User, error) {
	var user userdomain.User
	
	db := repository.DB.WithContext(ctx)
	result := db.First(&user, "username = ?", username)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err := result.Scan(&user)
		exception.CheckError(err.Error)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}	
}