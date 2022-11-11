package repository

import (
	"github.com/bimaagung/cafe-reservation/entity"
	"github.com/bimaagung/cafe-reservation/exception"
	"gorm.io/gorm"
)

func NewConnectDB(database *gorm.DB) MenuRepository {
	return &menuRepositoryImp{
		DB: database,
	} 
}

type menuRepositoryImp struct {
	DB *gorm.DB
}

func (repository *menuRepositoryImp) Add(menu entity.Menu) {  
	result := repository.DB.Create(&menu)
	exception.Error(result.Error)
}

func (repository *menuRepositoryImp) Delete(id string) {  
	result := repository.DB.Where("id = ?", id).Delete(&entity.Menu{})
	exception.Error(result.Error)
}

func (repository *menuRepositoryImp) GetByName(name string) entity.Menu {  
	var menu entity.Menu
	repository.DB.First(&menu, "name = ?", name)
	return menu
}

func (repository *menuRepositoryImp) GetById(id string) entity.Menu {  
	var menu entity.Menu
	repository.DB.First(&menu, "id = ?", id)
	return menu
}

func (repository *menuRepositoryImp) GetList() []entity.Menu {  
	var menu []entity.Menu
	result := repository.DB.Find(&menu)
	exception.Error(result.Error)
	return menu
}