package main

import (
	"log"

	"dadandev.com/dcbt/internal"
	"dadandev.com/dcbt/internal/config"
)

func main() {
	conf := config.Get()
	store := "dadan"
	server := internal.NewServer(*conf, store)
	log.Fatal(server.Start())
}
