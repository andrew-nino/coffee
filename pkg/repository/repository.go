package repository

import "github.com/jmoiron/sqlx"

type Authorisation interface {
}

type CoffeeList interface {
}

type CoffeeItem interface {
}

type Repository struct {
	Authorisation
	CoffeeList
	CoffeeItem
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{}
}
