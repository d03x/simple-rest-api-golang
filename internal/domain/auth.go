package domain

import (
	"context"

	"dadandev.com/dcbt/internal/dto"
)

type AuthService interface {
	Login(auth_dat dto.LoginReq) dto.AuthRes
	GetUser(c context.Context) ([]dto.UserRes, error)
}
