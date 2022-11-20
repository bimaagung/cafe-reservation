package repository

import (
	userdomain "github.com/bimaagung/cafe-reservation/user/domain"
	"github.com/gofiber/fiber/v2"
)

type UserRepository interface {
	Create(ctx *fiber.Ctx, user userdomain.User)
	GetById(ctx *fiber.Ctx, id string) (user userdomain.User)
}