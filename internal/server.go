package internal

import (
	"database/sql"
	"log"
	"runtime"

	"dadandev.com/dcbt/internal/api/auth"
	"dadandev.com/dcbt/internal/api/siswa"
	"dadandev.com/dcbt/internal/database"
	"dadandev.com/dcbt/internal/interfaces"
	"dadandev.com/dcbt/internal/middleware"
	"dadandev.com/dcbt/internal/repository"
	"dadandev.com/dcbt/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	config interfaces.Config
	db     *sql.DB
}

func NewServer(config interfaces.Config) *Server {
	return &Server{
		config: config,
	}
}
func (server *Server) InitalizeServer() {
	db := database.Connect(server.config.Database)
	server.db = db
}

// run server
func (server *Server) Start() error {
	runtime.GOMAXPROCS(2)
	app := fiber.New()
	api := app.Group("/api")
	authMiddleware := middleware.AuthenticationMiddleware()
	userRepository := repository.NewUserRepository(server.db)
	authService := services.NewAuth(userRepository)
	auth.NewAuth(api, authMiddleware, authService)
	siswa.NewSiswa(api, authMiddleware)
	go func() {
		err := app.Listen(server.config.AppConfig.Port)
		if err != nil {
			log.Fatalf("Server error %s", err.Error())
		}
	}()
	select {}

}
