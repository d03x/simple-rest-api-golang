package auth

import (
	"github.com/gofiber/fiber/v2"
)

type authApi struct {
}

func NewAuth(app *fiber.App) {
	ha := authApi{}
	app.Get("/login", ha.login)
}

func (app authApi) login(ctx *fiber.Ctx) error {
	return ctx.SendString("Oke login")
}
