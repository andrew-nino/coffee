package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
)

type ActionService struct {
	repo repository.CoffeeAction
}

func NewActionServise(repo repository.CoffeeAction) *ActionService {
	return &ActionService{repo: repo}
}

func (c *ActionService) GetActions() ([]coffee.Action, error) {
	return c.repo.GetActions()
}

func (c *ActionService) GetActionById(guid string) (coffee.Action, error) {
	return c.repo.GetActionById(guid)
}
