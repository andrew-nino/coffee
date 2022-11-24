package repository

import (
	"coffee-app"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user coffee.User) (int, error)
	GetUser(username, password string) (coffee.User, error)
}

type CoffeeList interface {
}

type CoffeeItem interface {
}

type Repository struct {
	Authorization
	CoffeeList
	CoffeeItem
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
