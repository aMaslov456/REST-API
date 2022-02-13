package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/aMaslov456/REST-API"
	"github.com/aMaslov456/REST-API/pkg/repository"
)

type AuthService struct {
	repo repository.Authorization
}

const salt = "sdfjnsdlkfndsfndsf"

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
