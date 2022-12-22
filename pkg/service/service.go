package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user coffee.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (string, error)
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

type CoffeeAction interface {
	GetActions() ([]coffee.Action, error)
	GetActionById(guid string) (coffee.Action, error)
}

type CoffeeDBUpdate interface {
	UpdateDB() (string, error)
	UpdatePoints(phone string, points float32) (coffee.User, error)
	UpdateUser(coffee.User) error
}

type Service struct {
	Authorization
	CoffeeList
	CoffeeItem
	CoffeeTypes
	CoffeeDBUpdate
	CoffeeAction
}

func NewService(repos *repository.Repository) *Service {

	return &Service{
		Authorization:  NewAuthService(repos.Authorization),
		CoffeeList:     NewAllCategoriesPostgres(repos.CoffeeList),
		CoffeeItem:     NewItemsById(repos.CoffeeItem),
		CoffeeTypes:    NewTypesService(repos.CoffeeTypes),
		CoffeeDBUpdate: NewUpdateService(repos.CoffeeDBUpdate),
		CoffeeAction:   NewActionServise(repos.CoffeeAction),
	}
}
