package usecase

import "github.com/bimaagung/cafe-reservation/menu/domain"

type MenuUseCase interface {
	GetList() (response []domain.MenuRes)
	GetById(id string) (response domain.MenuRes)
	Add(request domain.MenuReq, urlFile string) (response domain.MenuRes)
	Update(id string, request domain.MenuReq) (response domain.MenuRes)
	Delete(id string) bool
}
