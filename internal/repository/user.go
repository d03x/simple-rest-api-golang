package repository

import (
	"context"
	"database/sql"

	"dadandev.com/dcbt/internal/domain"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
)

type userRepository struct {
	db *goqu.Database
}

func NewUserRepository(db *sql.DB) domain.UserRepository {
	return &userRepository{
		db: goqu.Dialect("mysql").DB(db),
	}
}

func (u *userRepository) Save(ctx context.Context, user *domain.User) error {
	dataset := u.db.Insert("users").Rows(&user).Executor()
	_, err := dataset.ExecContext(ctx)
	return err
}
func (u *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	user := []domain.User{}
	dataset := u.db.From("users").Select(&domain.User{})
	err := dataset.ScanStructsContext(ctx, &user)
	return user, err
}
func (u *userRepository) FindByEmail(ctx context.Context, user domain.User) (domain.User, error) {
	users := domain.User{}
	_, err := u.db.From("users").Select(&domain.User{}).Where(goqu.Ex{
		"email": user.Email,
	}).ScanStructContext(ctx, &users)

	if err != nil {
		return domain.User{}, err
	}

	return users, err

}
