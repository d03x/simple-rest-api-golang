package auth

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"dadandev.com/dcbt/internal/domain"
	"dadandev.com/dcbt/internal/dto"
	"dadandev.com/dcbt/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	service domain.AuthService
}

func NewAuth(app fiber.Router, authHandler fiber.Handler, service domain.AuthService) {
	handler := authApi{
		service: service,
	}
	app.Get("/", handler.validateJwt)
	app.Get("/login", handler.login)
	app.Get("/users", authHandler, handler.getUser)
}
func (app authApi) validateJwt(ctx *fiber.Ctx) error {

	dat, err := utils.ValidateJwt("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhGFuQGdtYWlsLmNvbSIsImV4cCI6MTc0MTk2NjI0NiwiaWQiOiI2ODE4MmExMy1mZmU1LTExZWYtYjdlZi1mMGRlZjFiOWVmNzQiLCJuYW1lIjoiaGVucmEifQ.syAR9evrRfqF-QfhNHD7WEOhmjBXbQxB-qSFGIaYJPc")
	if err != nil {
		slog.ErrorContext(ctx.Context(), err.Error())
		return ctx.Status(http.StatusInternalServerError).JSON(dto.NewResponseMessage("Token is Invalid"))
	}
	return ctx.Status(http.StatusOK).JSON(dto.NewResponseMessage(dat.Email))
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
	loc := ctx.Locals("x-user")
	if loc != nil {
		fmt.Println(loc)
	}
	c, cancel := context.WithTimeout(ctx.Context(), 10+time.Second)
	defer cancel()
	users, err := app.service.GetUser(c)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(dto.NewResponseMessage(err.Error()))
	}
	return ctx.Status(fiber.StatusOK).JSON(dto.NewResponseData[[]dto.UserRes](users))
}
