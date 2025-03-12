package services

import (
	"fmt"

	"dadandev.com/dcbt/internal/domain"
)

type authService struct {
}

func NewAuth() domain.AuthService {
	return authService{}
}

func (as authService) Login() {
	fmt.Println("Ecxecute from login in services")
}
