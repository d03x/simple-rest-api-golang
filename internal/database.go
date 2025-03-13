package internal

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_golang")

	if err != nil {
		return nil, err
	}
	return db, nil
}
