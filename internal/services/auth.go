package services

import (
	"fmt"

	"dadandev.com/dcbt/internal/domain"
	"dadandev.com/dcbt/internal/dto"
)

type authService struct {
}

func NewAuth() domain.AuthService {
	return authService{}
}

func (as authService) Login(data dto.LoginReq) dto.AuthRes {
	var res = dto.AuthRes{}
	res.AccessToken = "3243225235"
	fmt.Println("Ecxecute from login in services")
	return res
}
