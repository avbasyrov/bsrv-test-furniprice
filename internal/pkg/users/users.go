package users

import (
	"context"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/dbcon"
	"github.com/avbasyrov/bsrv-test-furniprice/internal/pkg/models"
)

type Users struct {
	db *dbcon.Db
}

func New(db *dbcon.Db) *Users {
	return &Users{
		db: db,
	}
}

func (u *Users) Add(ctx context.Context, login string, password string) (models.User, error) {
	const query = "INSERT INTO public.users (login, password) VALUES ($1, $2) RETURNING id"
	lastInsertId := 0

	row := u.db.Sqlx.QueryRowContext(ctx, query, login, password)
	err := row.Scan(&lastInsertId)
	if err != nil {
		return models.User{}, err
	}

	return models.User{
		ID:    lastInsertId,
		Login: login,
		Name:  login,
	}, nil
}

func (u *Users) GetByLoginAndPassword(ctx context.Context, login string, password string) (models.User, error) {
	const query = "SELECT id, login, login AS name, " +
		"CASE WHEN admin IS true THEN 'admin' ELSE 'user' END AS role " +
		"FROM public.users " +
		"WHERE login LIKE $1 AND password LIKE $2"
	var user models.User
	err := u.db.Sqlx.GetContext(ctx, &user, query, login, password)
	return user, err
}
