package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
)

type UpdateService struct {
	repo repository.CoffeeDBUpdate
}

func NewUpdateService(repo repository.CoffeeDBUpdate) *UpdateService {
	return &UpdateService{repo: repo}
}

func (u *UpdateService) UpdateDB() (string, error) {
	return u.repo.UpdateDB()
}

func (u *UpdateService) UpdatePoints(phone string, points float32) (coffee.User, error) {

	phone = generatePasswordHash(phone)

	return u.repo.UpdatePoints(phone, points)
}

func (u *UpdateService) UpdateUser(user coffee.User) error {

	user.Phone = generatePasswordHash(user.Phone)

	return u.repo.UpdateUser(user)
}
