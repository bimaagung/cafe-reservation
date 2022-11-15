package usecase

import (
	"github.com/bimaagung/cafe-reservation/models/domain"
)

type MenuUseCase interface {
	GetList() (response []domain.Menu)
	Add(request domain.Menu) (response domain.Menu)
	Delete(id string) bool
}
