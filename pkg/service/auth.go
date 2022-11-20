package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const solt = "jbwdvkjbvjwnvlwk"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthservice(repo repository.Authorization) *AuthService {
	
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user coffee.User) (int, error) {

	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {

	hach := sha1.New()
	hach.Write([]byte(password))

	return fmt.Sprintf("%x", hach.Sum([]byte(solt)))
}
