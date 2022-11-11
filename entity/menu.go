package entity

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	Id    string
	Name  string
	Price int64
	Stock int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}