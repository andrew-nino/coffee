package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user coffee.User) (int, error)
	GenerateToken(username, password string) (string, error)
}

type CoffeeList interface {
}

type CoffeeItem interface {
}

type Service struct {
	Authorization
	CoffeeList
	CoffeeItem
}

func NewService(repos *repository.Repository) *Service {

	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
