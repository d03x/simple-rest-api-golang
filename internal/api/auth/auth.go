package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"dadandev.com/dcbt/internal/domain"
	"dadandev.com/dcbt/internal/dto"
	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	service domain.AuthService
}

func NewAuth(app fiber.Router, service domain.AuthService) {
	handler := authApi{
		service: service,
	}
	app.Get("/login", handler.login)
	app.Get("/users", handler.getUser)
}

func (app authApi) login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 3*time.Second)
	defer cancel()
	data := dto.LoginReq{}
	data.Email = "dadan@gmail.com"
	data.Password = "Sumedang"
	res, err := app.service.Login(c, data)
	if err != nil {
		if errors.Is(err, domain.InvalidCredential) {
			return ctx.Status(fiber.StatusUnauthorized).JSON(dto.NewResponseMessage("Invalid credentials tolong cek kredensial anda atau pake user lain"))
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.NewResponseMessage("An internal server error has occurred. Please try again later. If the issue persists, contact support for further assistance. We apologize for any inconvenience."))
	}
	return ctx.Status(http.StatusOK).JSON(dto.NewResponseData[dto.AuthRes](res))
}

func (app authApi) getUser(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10+time.Second)
	defer cancel()
	users, err := app.service.GetUser(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.NewResponseMessage(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.NewResponseData[[]dto.UserRes](users))
}
