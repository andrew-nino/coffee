package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
)

type TypeService struct {
	repo repository.CoffeeTypes
}

func NewTypesService(repo repository.CoffeeTypes) *TypeService {
	return &TypeService{repo: repo}
}

func (c *TypeService) GetTypes(item string) ([]coffee.Type, error) {
	return c.repo.GetTypes(item)
}
