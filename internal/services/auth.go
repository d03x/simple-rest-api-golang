package services

import (
	"database/sql"
	"fmt"
	"log"

	"dadandev.com/dcbt/internal/domain"
	"dadandev.com/dcbt/internal/dto"
)

type authService struct {
	db *sql.DB
}

func NewAuth(db *sql.DB) domain.AuthService {
	return authService{
		db: db,
	}
}

func (as authService) Login(data dto.LoginReq) dto.AuthRes {
	var res = dto.AuthRes{}
	res.AccessToken = "3243225235"

	return res
}

func (as authService) GetUser() []dto.UserRes {

	rows, err := as.db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer rows.Close()

	var userdata []dto.UserRes

	for rows.Next() {
		var each = dto.UserRes{}
		err = rows.Scan(&each.Id, &each.Name, &each.Email)
		if err != nil {
			fmt.Println(err.Error())
		}
		userdata = append(userdata, each)
	}
	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(userdata)
	return userdata
}
