package repository

import (
	"context"

	"github.com/bimaagung/cafe-reservation/menu/domain"
)

type MenuRepository interface {
	Add(ctx context.Context, menu *domain.Menu) (string, error)
	GetList(ctx context.Context) (menu []domain.Menu, err error)
	Delete(ctx context.Context, id string) error
	GetByName(ctx context.Context, name string) (menu domain.Menu, err error)
	GetById(ctx context.Context, id string) (menu domain.Menu, err error)
	Update(ctx context.Context, id string, menu *domain.Menu) (string, error)
}