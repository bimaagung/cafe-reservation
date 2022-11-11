package usecase

import "github.com/bimaagung/cafe-reservation/model"

type MenuUseCase interface {
	GetList()(response []model.Menu)
	Add(request model.Menu)(response model.Menu)
	Delete(id string) bool
}