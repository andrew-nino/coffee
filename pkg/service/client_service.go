package service

import "coffee-app/pkg/repository"

type ClientService struct {
	repo repository.CoffeeClient
}

func NewClientServece(repo repository.CoffeeClient) *ClientService {
	return &ClientService{repo: repo}
}

func (c *ClientService) GetBalance(id int) (float32, error) {
	return c.repo.GetBalance(id)
}
