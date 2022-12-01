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
	GetCategories(category string) ([]coffee.Category, error)
}

type CoffeeItem interface {
	GetItemsById(cat string) ([]coffee.Item, error)
	GetItems() ([]coffee.Item, error)
}

type CoffeeTypes interface {
	GetTypes(item string) ([]coffee.Type, error)
}

type Repository struct {
	Authorization
	CoffeeList
	CoffeeItem
	CoffeeTypes
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Authorization: NewAuthPostgres(db),
		CoffeeList:    NewAllCategoriesPostgres(db),
		CoffeeItem:    NewItemsostgres(db),
		CoffeeTypes:   NewTypesostgres(db),
	}
}
