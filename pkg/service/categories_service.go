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

func (c *CategoriesService) GetALLCategories() ([]coffee.Categories, error) {
	return c.repo.GetALLCategories()
}
