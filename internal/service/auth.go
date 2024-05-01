package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"to-do-checklist/internal/config"
	"to-do-checklist/internal/domain"
	"to-do-checklist/internal/repository"
)

type AuthService struct {
	cfg  *config.AuthConfig
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization, authConfig *config.AuthConfig) *AuthService {
	return &AuthService{repo: repo, cfg: authConfig}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.PasswordHashingSalt)))
}

type AuthError error

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func (s *AuthService) GenerateToken(authInfo domain.SignIn) (string, error) {
	user, err := s.repo.GetUser(authInfo.Username)
	if err != nil {
		return "", err
	}
	if user.Password != s.generatePasswordHash(authInfo.Password) {
		return "", AuthError(errors.New("password does not match"))
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(s.cfg.TokenTTL)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString([]byte(s.cfg.JWTSigningKey))
}
