package usecase

import (
	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/gofiber/fiber/v2"
)

type UserUseCase interface {
	Create(ctx *fiber.Ctx, request userdomain.UserReq) userdomain.UserRes
}