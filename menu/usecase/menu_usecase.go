package usecase

import (
	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/gofiber/fiber/v2"
)

type MenuUseCase interface {
	GetList(ctx *fiber.Ctx) (response []domain.MenuRes)
	GetById(ctx *fiber.Ctx, id string) (response domain.MenuRes)
	Add(ctx *fiber.Ctx, request domain.MenuReq) (response domain.MenuRes)
	Update(ctx *fiber.Ctx, id string, request domain.MenuReq) (response domain.MenuRes)
	Delete(ctx *fiber.Ctx, id string) bool
}
