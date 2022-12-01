package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
)

type CategoriesService struct {
	repo repository.CoffeeList
}

func NewAllCategoriesPostgres(repo repository.CoffeeList) *CategoriesService {
	return &CategoriesService{repo: repo}
}

func (c *CategoriesService) GetCategories(category string) ([]coffee.Category, error) {
	return c.repo.GetCategories(category)
}
