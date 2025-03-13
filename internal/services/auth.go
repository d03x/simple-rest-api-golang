package services

import (
	"context"

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

func (as authService) Login(data dto.LoginReq) dto.AuthRes {
	var res = dto.AuthRes{}
	res.AccessToken = "3243225235"

	return res
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
