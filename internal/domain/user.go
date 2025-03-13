package domain

import (
	"context"
)

type User struct {
	Id         string `db:"id"`
	Name       string `db:"name"`
	Email      string `db:"email"`
	DeviceId   string `db:"device_id"`
	LastActive string `db:"last_active"`
	Password   string `db:"password"`
}

type UserRepository interface {
	Save(ctx context.Context, user *User) error
	GetAll(ctx context.Context) ([]User, error)
}
