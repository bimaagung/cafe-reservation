package menurepository

import (
	"github.com/bimaagung/cafe-reservation/models/domain"
)

type MenuRepository interface {
	Add(menu domain.Menu)
	GetList() []domain.Menu
	Delete(id string)
	GetByName(name string) domain.Menu
	GetById(id string) domain.Menu
	Update(id string, menu domain.Menu)
}