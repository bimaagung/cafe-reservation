package domain

import (
	"time"

	"gorm.io/gorm"
)

// Menu
type Menu struct {
	Id    string `json:"id"`
	Name  string `json:"name" validate:"required"` 
	Price int64 `json:"price" validate:"required"`
	Stock int `json:"stock" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type MenuUseCase interface {
	GetList()(response []Menu)
	Add(request Menu)(response Menu)
	Delete(id string) bool
}

type MenuRepository interface {
	Add(menu Menu)
	GetList()[]Menu
	Delete(id string)
	GetByName(name string) Menu
	GetById(id string) Menu
}