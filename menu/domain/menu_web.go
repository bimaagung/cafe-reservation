package domain

import (
	"mime/multipart"
	"time"
)

type MenuReq struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"` 
	Price int64 `json:"price,omitempty"`
	Stock int `json:"stock,omitempty"`
	File *multipart.FileHeader
}

type MenuRes struct {
	Id    string `json:"id"`
	Name  string `json:"name" validate:"required"` 
	Price int64 `json:"price" validate:"required"`
	Stock int `json:"stock" validate:"required"`
	Url  string `json:"url" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


