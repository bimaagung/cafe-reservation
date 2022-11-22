package repository

import (
	"context"

	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user userdomain.User) 
	GetByUsername(ctx context.Context, username string) (user userdomain.User, err error)
}