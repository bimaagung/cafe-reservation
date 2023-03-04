package domain

import (
	"context"
	"mime/multipart"
	"time"

	"gorm.io/gorm"
)

// Menu
type Menu struct {
	Id    string 
	Name  string  
	Price int64 
	Stock int 
	Url string
	CreatedAt time.Time 
	UpdatedAt time.Time 
	DeletedAt gorm.DeletedAt 
}

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

type MenuUseCase interface {
	GetList(ctx context.Context) (response []MenuRes, err error)
	GetById(ctx context.Context, id string) (response MenuRes, err error)
	Add(ctx context.Context, request *MenuReq) (response MenuRes, err error)
	Update(ctx context.Context, id string, request *MenuReq) (response MenuRes, err error)
	Delete(ctx context.Context, id string) (bool, error)
}


type MinioRepository interface {
	Upload(file *multipart.FileHeader, bucketName string, objectName string) error
}

type MenuRepository interface {
	Add(ctx context.Context, menu *Menu) (string, error)
	GetList(ctx context.Context) (menu []Menu, err error)
	Delete(ctx context.Context, id string) error
	GetByName(ctx context.Context, name string) (menu Menu, err error)
	GetById(ctx context.Context, id string) (menu Menu, err error)
	Update(ctx context.Context, id string, menu *Menu) (string, error)
}

type MenuRepositoryRedis interface {
	Set(menu interface{}) (string, error)
	Get() (menu []Menu, err error)
	Delete() error
}