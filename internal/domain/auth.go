package domain

import "github.com/gofiber/fiber/v2"

type AuthService interface {
	Login(ctx *fiber.Ctx) error
}
