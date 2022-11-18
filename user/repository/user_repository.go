package repository

import (
	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
)

type UserRepository interface {
	Create(user userdomain.User)
	GetById(id string) (user userdomain.User)
}