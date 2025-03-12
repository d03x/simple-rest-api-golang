package internal

import (
	"dadandev.com/dcbt/internal/api/auth"
	"dadandev.com/dcbt/internal/interfaces"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config interfaces.Config
	store  string
}

func NewServer(config interfaces.Config, store string) *Server {
	return &Server{
		config: config,
		store:  store,
	}
}

func (s *Server) Start() error {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello world")
	})
	auth.NewAuth(app)
	return app.Listen(s.config.AppConfig.Port)
}
