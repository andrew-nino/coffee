package service

import "coffee-app/pkg/repository"

type UpdateService struct {
	repo repository.CoffeeDBUpdate
}

func NewUpdateService(repo repository.CoffeeDBUpdate) *UpdateService {
	return &UpdateService{repo: repo}
}

func (u *UpdateService) UpdateDB() (string, error) {
	return u.repo.UpdateDB()
}

func (u *UpdateService) UpdatePoints(phone string, points float32) (float32, error) {

	phone = generatePasswordHash(phone)

	return u.repo.UpdatePoints(phone, points)
}
