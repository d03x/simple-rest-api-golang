package middleware

import (
	"strings"

	"dadandev.com/dcbt/internal/dto"
	"dadandev.com/dcbt/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func AuthenticationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := strings.Split(c.Get("Authorization"), " ")

		if len(token) < 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.NewResponseMessage("Maaf token tidak ada tolong cobalagi"))
		}

		data, err := utils.ValidateJwt(token[1])
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.NewResponseMessage("Maaf token tidak ada tolong cobalagi"))
		}
		c.Locals("x-user", data)

		return c.Next()
	}
}
