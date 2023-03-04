package postgres

import (
	"context"
	"errors"

	"github.com/bimaagung/cafe-reservation/domain"
	"gorm.io/gorm"
)

func NewConnectDB(database *gorm.DB) domain.MenuRepository {
	return &postgresMenuRepository{
		DB: database,
	} 
}

type postgresMenuRepository struct {
	DB *gorm.DB
}

func (repository *postgresMenuRepository) Add(ctx context.Context, menu *domain.Menu) (string, error) {  
	db := repository.DB.WithContext(ctx)
	err := db.Create(&menu).Error
	if err != nil {
		return "", err
	}

	return menu.Id, nil
}

func (repository *postgresMenuRepository) Delete(ctx context.Context, id string) error {  
	db := repository.DB.WithContext(ctx)
	err := db.Where("id = ?", id).Delete(&domain.Menu{}).Error
	if err != nil {
		return  err
	}

	return  nil
}

func (repository *postgresMenuRepository) GetByName(ctx context.Context, name string) (menu domain.Menu, err error) {  
	db := repository.DB.WithContext(ctx)
	err = db.First(&menu, "name = ?", name).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return menu, err  
	}

	return menu, nil
}

func (repository *postgresMenuRepository) GetById(ctx context.Context, id string) (menu domain.Menu, err error){  
	db := repository.DB.WithContext(ctx)
	err = db.First(&menu, "id = ?", id).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return menu, err  
	}
	return menu, nil
}

func (repository *postgresMenuRepository) GetList(ctx context.Context) (menu []domain.Menu, err error) {  
	db := repository.DB.WithContext(ctx)
	err = db.Find(&menu).Error

	if err != nil {
		return nil, err
	}

	return menu, nil
}

func (repository *postgresMenuRepository) Update(ctx context.Context, id string, menu *domain.Menu) (string, error) {
	db := repository.DB.WithContext(ctx)
	err := db.Where("id = ?", id).Updates(&menu).Error
	
	if err != nil {
		return menu.Id, err
	}

	return menu.Id, nil
}