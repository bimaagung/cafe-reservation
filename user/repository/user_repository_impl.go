package repository

import (
	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/bimaagung/cafe-reservation/utils/exception"
	"github.com/gofiber/fiber/v2"
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

func (repository *postgresUserRepository) Create(ctx *fiber.Ctx, user userdomain.User) {

	db := repository.DB.WithContext(ctx.Context())
	result := db.Create(&user)
	exception.Error(result.Error)
	
}

func (repository *postgresUserRepository) GetById(ctx *fiber.Ctx, id string)(user userdomain.User) {
	db := repository.DB.WithContext(ctx.Context())
	db.First(&user, "id = ?", id)
	return user
}