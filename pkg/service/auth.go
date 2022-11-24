package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	solt        = "jbwdvkjbvjwnvlwk"
	signingtKey = "kmvuoivpw;mfnvojnweof"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {

	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user coffee.User) (int, error) {

	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {

	user, err := s.repo.GetUser(username, generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},

		user.Id,
	})

	return token.SignedString([]byte(signingtKey))
}

func generatePasswordHash(password string) string {

	hach := sha1.New()
	hach.Write([]byte(password))

	return fmt.Sprintf("%x", hach.Sum([]byte(solt)))
}
