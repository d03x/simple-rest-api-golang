package siswa

import "github.com/gofiber/fiber/v2"

type siswaApi struct {
}

func NewSiswa(router fiber.Router) {
	c := siswaApi{}
	siswa := router.Group("siswa")
	siswa.Get("get", c.fetch)
}

func (c siswaApi) fetch(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON("[]")
}
