package domain

import (
	"context"

	"dadandev.com/dcbt/internal/dto"
)

type AuthService interface {
	Login(ctx context.Context, auth_dat dto.LoginReq) (dto.AuthRes, error)
	GetUser(c context.Context) ([]dto.UserRes, error)
}
