package postgres

import (
	"context"

	"github.com/bimaagung/cafe-reservation/domain"
	"gorm.io/gorm"
)


func NewUserRepository(database *gorm.DB) domain.UserRepository {
	return &postgresUserRepository {
		DB: database,
	}
}

type postgresUserRepository struct {
	DB *gorm.DB
}

func (repository *postgresUserRepository) Create(ctx context.Context, user *domain.User) (string, error) {
	db := repository.DB.WithContext(ctx)
	err := db.Create(&user).Error

	if err != nil {
		return "", err
	}

	return user.Id, nil
}


func (repository *postgresUserRepository) GetByUsername(ctx context.Context, username string)(user domain.User, err error) {
	
	db := repository.DB.WithContext(ctx)
	result := db.First(&user, "username = ?", username)
	return user, result.Error
}
