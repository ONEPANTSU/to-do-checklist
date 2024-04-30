package service

import (
	"crypto/sha1"
	"fmt"
	"to-do-checklist/internal/domain"
	"to-do-checklist/internal/repository"
)

const salt = "s2dfd8s23ifnds9if234asnwjovwvdemv32e42rfce2ds"

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
