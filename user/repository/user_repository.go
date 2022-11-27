package repository

import (
	"context"

	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *userdomain.User) (string, error) 
	GetByUsername(ctx context.Context, username string) (userdomain.User, error)
}