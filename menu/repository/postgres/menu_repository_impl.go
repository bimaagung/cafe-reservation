package repository

import (
	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/bimaagung/cafe-reservation/utils/exception"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewConnectDB(database *gorm.DB) MenuRepository {
	return &postgresMenuRepository{
		DB: database,
	} 
}

type postgresMenuRepository struct {
	DB *gorm.DB
}

func (repository *postgresMenuRepository) Add(ctx *fiber.Ctx,menu domain.Menu) {  
	db := repository.DB.WithContext(ctx.Context())
	result := db.Create(&menu)
	exception.CheckError(result.Error)
}

func (repository *postgresMenuRepository) Delete(ctx *fiber.Ctx, id string) {  
	db := repository.DB.WithContext(ctx.Context())
	result := db.Where("id = ?", id).Delete(&domain.Menu{})
	exception.CheckError(result.Error)
}

func (repository *postgresMenuRepository) GetByName(ctx *fiber.Ctx, name string) domain.Menu {  
	var menu domain.Menu

	db := repository.DB.WithContext(ctx.Context())
	db.First(&menu, "name = ?", name)
	return menu
}

func (repository *postgresMenuRepository) GetById(ctx *fiber.Ctx, id string) domain.Menu {  
	var menu domain.Menu

	db := repository.DB.WithContext(ctx.Context())
	db.First(&menu, "id = ?", id)
	return menu
}

func (repository *postgresMenuRepository) GetList(ctx *fiber.Ctx) []domain.Menu {  
	var menu []domain.Menu

	db := repository.DB.WithContext(ctx.Context())
	result := db.Find(&menu)

	exception.CheckError(result.Error)
	return menu
}

func (repository *postgresMenuRepository) Update(ctx *fiber.Ctx, id string, menu domain.Menu) {
	db := repository.DB.WithContext(ctx.Context())
	result := db.Where("id = ?", id).Updates(&menu)
	exception.CheckError(result.Error)
}