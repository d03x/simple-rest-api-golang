package domain

import "dadandev.com/dcbt/internal/dto"

type AuthService interface {
	Login(auth_dat dto.LoginReq) dto.AuthRes
}
