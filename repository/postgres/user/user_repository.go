package userrepository

import "github.com/bimaagung/cafe-reservation/models/domain"

type UserRepository interface {
	Create(user domain.User)
	GetById(id string)(user domain.User)
}