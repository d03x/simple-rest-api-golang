package internal

import (
	"database/sql"
	"fmt"
	"log"
	"runtime"
	"time"

	"dadandev.com/dcbt/internal/api/auth"
	"dadandev.com/dcbt/internal/database"
	"dadandev.com/dcbt/internal/interfaces"
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
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}
	server.db = db
}

// run server
func (server *Server) Start() error {
	runtime.GOMAXPROCS(2)
	app := fiber.New()
	defer server.db.Close()
	//for services
	authService := services.NewAuth(server.db)
	auth.NewAuth(app, authService)
	go func() {
		err := app.Listen(server.config.AppConfig.Port)
		if err != nil {
			log.Fatalf("Server error %s", err.Error())
		}
	}()
	go func() {
		status := "Server berjalan"
		_, err := server.db.Exec("INSERT INTO logs (time,log) VALUES(?,?)", time.Now(), status)
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	fmt.Print("Tidak menunggu Listen")
	select {}

}
