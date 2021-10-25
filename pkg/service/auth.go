package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/Red-373/todo"
	"github.com/Red-373/todo/pkg/repository"
)

const salt = "sjadnasjdnjaandjandjansjdnjn1j3njs"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generetePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generetePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
