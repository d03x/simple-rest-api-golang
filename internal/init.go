package internal

import (
	"log"

	"dadandev.com/dcbt/internal/config"
)

// INITIALIZE APP
func AppInit() {
	config := config.Get()
	server := NewServer(*config)
	//initialize seperti koneksi database dan lain lain
	server.InitalizeServer()
	//jalankan server
	log.Fatal(server.Start())
}
