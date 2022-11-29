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
	GetALLCategories() ([]coffee.Categories, error)
}

type CoffeeItem interface {
	GetItemsById(category string) ([]coffee.Items, error)
}

type Service struct {
	Authorization
	CoffeeList
	CoffeeItem
}

func NewService(repos *repository.Repository) *Service {

	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		CoffeeList:    NewAllCategoriesPostgres(repos.CoffeeList),
		CoffeeItem:    NewItemsById(repos.CoffeeItem),
	}
}
