package services

import (
	"context"
	"fmt"

	"dadandev.com/dcbt/internal/domain"
	"dadandev.com/dcbt/internal/dto"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
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

	fmt.Println(user)
	var res = dto.AuthRes{
		AccessToken: user.Email,
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
