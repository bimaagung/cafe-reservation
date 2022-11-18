package usecase

import userdomain "github.com/bimaagung/cafe-reservation/user/domain"

type UserUseCase interface {
	Create(request userdomain.UserReq) userdomain.UserRes
}