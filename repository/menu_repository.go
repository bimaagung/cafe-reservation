package repository

import "github.com/bimaagung/cafe-reservation/entity"

type MenuRepository interface {
	Add(menu entity.Menu)
	GetList()[]entity.Menu
	Delete(id string)
	GetByName(name string) entity.Menu
	GetById(id string) entity.Menu
}