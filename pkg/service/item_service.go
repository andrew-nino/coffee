package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
)

type ItemsService struct {
	repo repository.CoffeeItem
}

func NewItemsById(repo repository.CoffeeItem) *ItemsService {
	return &ItemsService{repo: repo}
}

func (c *ItemsService) GetItemsById(categories string) ([]coffee.Item, error) {
	return c.repo.GetItemsById(categories)
}
func (c *ItemsService) GetItems() ([]coffee.Item, error) {
	return c.repo.GetItems()
}