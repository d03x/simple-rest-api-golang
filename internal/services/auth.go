package services

import (
	"context"
	"time"

	"dadandev.com/dcbt/internal/domain"
	"dadandev.com/dcbt/internal/dto"
	"dadandev.com/dcbt/internal/utils"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/golang-jwt/jwt/v5"
)

type authService struct {
	repository domain.UserRepository
}

func NewAuth(repository domain.UserRepository) domain.AuthService {
	return authService{
		repository: repository,
	}
}

func (as authService) Login(ctx context.Context, data dto.LoginReq) (dto.AuthRes, error) {

	user, err := as.repository.FindByEmail(ctx, domain.User{
		Email: data.Email,
	})
	if err != nil {
		return dto.AuthRes{}, err
	}
	if user.Email == "" {
		return dto.AuthRes{}, domain.InvalidCredential
	}
	token, err := utils.CreateJwtToken(jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
		"email": user.Email,
		"name":  user.Name,
		"id":    user.Id,
	})
	if err != nil {
		return dto.AuthRes{}, err
	}
	var res = dto.AuthRes{
		AccessToken: token,
	}
	return res, err
}

func (as authService) GetUser(ctx context.Context) ([]dto.UserRes, error) {
	res := []dto.UserRes{}
	user, err := as.repository.GetAll(ctx)

	if err != nil {
		return []dto.UserRes{}, err
	}

	for _, data := range user {
		res = append(res, dto.UserRes{
			Email:      data.Email,
			Id:         data.Id,
			Name:       data.Name,
			DeviceId:   data.DeviceId,
			LastActive: data.LastActive,
		})
	}

	return res, err
}
