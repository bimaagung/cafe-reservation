package usecase

import (
	"context"

	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
)

type UserUseCase interface {
	Create(ctx context.Context, request userdomain.UserReq) (userdomain.UserRes, error)
}