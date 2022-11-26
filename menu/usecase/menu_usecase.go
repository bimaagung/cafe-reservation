package usecase

import (
	"context"

	"github.com/bimaagung/cafe-reservation/menu/domain"
)

type MenuUseCase interface {
	GetList(ctx context.Context) (response []domain.MenuRes, err error)
	GetById(ctx context.Context, id string) (response domain.MenuRes, err error)
	Add(ctx context.Context, request *domain.MenuReq) (response domain.MenuRes, err error)
	Update(ctx context.Context, id string, request *domain.MenuReq) (response domain.MenuRes, err error)
	Delete(ctx context.Context, id string) (bool, error)
}
