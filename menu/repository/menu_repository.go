package repository

import (
	"github.com/bimaagung/cafe-reservation/menu/domain"
	"github.com/gofiber/fiber/v2"
)

type MenuRepository interface {
	Add(ctx *fiber.Ctx, menu domain.Menu)
	GetList(ctx *fiber.Ctx) []domain.Menu
	Delete(ctx *fiber.Ctx, id string)
	GetByName(ctx *fiber.Ctx,name string) domain.Menu
	GetById(ctx *fiber.Ctx,id string) domain.Menu
	Update(ctx *fiber.Ctx, id string, menu domain.Menu)
}