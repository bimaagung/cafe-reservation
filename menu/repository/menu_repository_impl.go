package repository

import (
	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/bimaagung/cafe-reservation/utils/exception"
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

func (repository *postgresMenuRepository) Add(menu domain.Menu) {  
	result := repository.DB.Create(&menu)
	exception.Error(result.Error)
}

func (repository *postgresMenuRepository) Delete(id string) {  
	result := repository.DB.Where("id = ?", id).Delete(&domain.Menu{})
	exception.Error(result.Error)
}

func (repository *postgresMenuRepository) GetByName(name string) domain.Menu {  
	var menu domain.Menu
	repository.DB.First(&menu, "name = ?", name)
	return menu
}

func (repository *postgresMenuRepository) GetById(id string) domain.Menu {  
	var menu domain.Menu
	repository.DB.First(&menu, "id = ?", id)
	return menu
}

func (repository *postgresMenuRepository) GetList() []domain.Menu {  
	var menu []domain.Menu
	result := repository.DB.Find(&menu)
	exception.Error(result.Error)
	return menu
}

func (repository *postgresMenuRepository) Update(id string, menu domain.Menu) {
	result := repository.DB.Where("id = ?", id).Updates(&menu)
	exception.Error(result.Error)
}