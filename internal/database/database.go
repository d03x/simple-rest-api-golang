package database

import (
	"database/sql"
	"fmt"
	"log"

	"dadandev.com/dcbt/internal/interfaces"
	_ "github.com/go-sql-driver/mysql"
)

func Connect(dbConf interfaces.Database) *sql.DB {
	database_url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Database)
	db, err := sql.Open("mysql", database_url)
	if err != nil {
		log.Fatal("Gagal koneksi kedatabase: ", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Gagal koneksi kedatabase: ", err.Error())
	}
	return db
}
