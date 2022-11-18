package domain

import (
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