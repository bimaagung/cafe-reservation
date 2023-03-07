package domain

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserReq struct {
	Name           string `json:"name"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RetypePassword string `json:"retype_password"`
}

type UserRes struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type UserUseCase interface {
	Create(ctx context.Context, request *UserReq) (UserRes, error)
}

type UserRepository interface {
	Create(ctx context.Context, user *User) (string, error) 
	GetByUsername(ctx context.Context, username string) (User, error)
}
