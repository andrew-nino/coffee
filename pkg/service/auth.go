package service

import (
	"coffee-app"
	"coffee-app/pkg/repository"
	"crypto/sha1"
	"errors"
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
	Item string `json:"item"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {

	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user coffee.User) (int, error) {

	user.Phone = generatePasswordHash(user.Phone)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(phoneCode, phone string) (string, error) {

	user, err := s.repo.GetUser(phoneCode, generatePasswordHash(phone))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
		},

		fmt.Sprint(user.Id),
	})

	return token.SignedString([]byte(signingtKey))
}

func (a *AuthService) ParseToken(accesToken string) (string, error) {

	token, err := jwt.ParseWithClaims(accesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingtKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Item, nil
}

func generatePasswordHash(password string) string {

	hach := sha1.New()
	hach.Write([]byte(password))

	return fmt.Sprintf("%x", hach.Sum([]byte(solt)))
}
