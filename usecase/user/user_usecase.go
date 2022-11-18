package userusecase

import (
	"github.com/bimaagung/cafe-reservation/models/web"
)

type UserUseCase interface {
	Create(request web.UserReq)(response web.UserRes)
}