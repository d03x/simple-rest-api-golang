package internal

import (
	"log"

	"dadandev.com/dcbt/internal/api/auth"
	"dadandev.com/dcbt/internal/interfaces"
	"dadandev.com/dcbt/internal/services"
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

// run server
func (s *Server) Start() error {
	db, err := Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	app := fiber.New()
	//for services
	authService := services.NewAuth(db)
	auth.NewAuth(app, authService)
	return app.Listen(s.config.AppConfig.Port)
}
