package config

import (
	"flag"
	"log"
	"os"

	"dadandev.com/dcbt/internal/interfaces"
	"github.com/joho/godotenv"
)

func Get() *interfaces.Config {

	flag_file := flag.String("env", "", "file .env location path absolute")
	flag.Parse()
	var err error
	if *flag_file != "" {
		err = godotenv.Load(*flag_file)
	} else {
		err = godotenv.Load()
	}
	if err != nil {
		log.Fatal(err.Error())
	}
	return &interfaces.Config{
		Database: interfaces.Database{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Database: os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
		Storage: interfaces.Storage{
			BackUpPath: os.Getenv("BACKUP_PATH"),
			UploadPath: os.Getenv("BACKUP_PATH"),
		},
		AppConfig: interfaces.AppConfig{
			Host: os.Getenv("HOST"),
			Port: os.Getenv("PORT"),
		},
	}
}
