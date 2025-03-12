package auth

import (
	"dadandev.com/dcbt/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	service domain.AuthService
}

func NewAuth(app *fiber.App, service domain.AuthService) {
	handler := authApi{
		service: service,
	}
	app.Get("/login", handler.login)
}

func (app authApi) login(ctx *fiber.Ctx) error {
	app.service.Login()
	return ctx.SendString("Oke login")
}
