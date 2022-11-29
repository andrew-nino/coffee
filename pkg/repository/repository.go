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
	GetALLCategories() ([]coffee.Categories, error)
}

type CoffeeItem interface {
	GetItemsById(cat string) ([]coffee.Items, error)
}

type Repository struct {
	Authorization
	CoffeeList
	CoffeeItem
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
		CoffeeList:    NewAllCategoriesPostgres(db),
		CoffeeItem:    NewItemsostgres(db),
	}
}
