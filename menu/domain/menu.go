package domain

import (
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


