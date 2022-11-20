package repository

import "github.com/bimaagung/cafe-reservation/menu/domain"

type MenuRepositoryRedis interface {
	Set(menu interface{}) (string, error)
	Get() (menu []domain.Menu, err error)
	Delete() error
}