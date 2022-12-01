package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user coffee.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type CoffeeList interface {
	GetCategories(category string) ([]coffee.Category, error)
}

type CoffeeItem interface {
	GetItemsById(category string) ([]coffee.Item, error)
	GetItems() ([]coffee.Item, error)
}

type CoffeeTypes interface {
	GetTypes(item string) ([]coffee.Type, error)
}

type Service struct {
	Authorization
	CoffeeList
	CoffeeItem
	CoffeeTypes
}

func NewService(repos *repository.Repository) *Service {

	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CoffeeList:    NewAllCategoriesPostgres(repos.CoffeeList),
		CoffeeItem:    NewItemsById(repos.CoffeeItem),
		CoffeeTypes:   NewTypesService(repos.CoffeeTypes),
	}
}
