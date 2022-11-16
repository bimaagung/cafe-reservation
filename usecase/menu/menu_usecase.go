package usecase

import (
	"github.com/bimaagung/cafe-reservation/models/web"
)

type MenuUseCase interface {
	GetList() (response []web.MenuRes)
	GetById(id string) (response web.MenuRes)
	Add(request web.MenuReq, urlFile string) (response web.MenuRes)
	Update(id string, request web.MenuReq) (response web.MenuRes)
	Delete(id string) bool
}
