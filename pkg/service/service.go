package service

import "coffee-app/pkg/repository"

type Authorisation interface {
}

type CoffeeList interface {
}

type CoffeeItem interface {
}

type Service struct {
	Authorisation
	CoffeeList
	CoffeeItem
}

func NewService(repos *repository.Repository) *Service {

	return &Service{}
}
