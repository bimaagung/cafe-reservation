package web

import (
	"time"
)

type MenuReq struct {
	Id    string `json:"id"`
	Name  string `json:"name" validate:"required"` 
	Price int64 `json:"price" validate:"required"`
	Stock int `json:"stock" validate:"required"`
}

type MenuRes struct {
	Id    string `json:"id"`
	Name  string `json:"name" validate:"required"` 
	Price int64 `json:"price" validate:"required"`
	Stock int `json:"stock" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


