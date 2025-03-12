package auth

import (
	"net/http"

	"dadandev.com/dcbt/internal/domain"
	"dadandev.com/dcbt/internal/dto"
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
	data := dto.LoginReq{}
	data.Email = "dadan@gmail.com"
	data.Password = "Sumedang"
	res := app.service.Login(data)
	return ctx.Status(http.StatusOK).JSON(dto.NewResponseData[dto.AuthRes](res))
}
