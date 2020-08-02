package models

import "context"

type User struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Role  string `json:"role"`
}

type UsersRepository interface {
	Add(ctx context.Context, login string, password string) (User, error)
	GetByLoginAndPassword(ctx context.Context, login string, password string) (User, error)
}
