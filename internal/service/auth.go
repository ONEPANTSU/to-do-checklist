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
	cfg      *config.AuthConfig
	authRepo repository.Authorization
}

func newAuthService(authRepo repository.Authorization, authConfig *config.AuthConfig) *AuthService {
	return &AuthService{authRepo: authRepo, cfg: authConfig}
}

func (s *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.authRepo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.cfg.PasswordHashingSalt)))
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func (s *AuthService) GenerateToken(authInfo domain.SignIn) (string, error) {
	user, err := s.authRepo.GetUser(authInfo.Username)
	if err != nil {
		return "", err
	}
	if user.Password != s.generatePasswordHash(authInfo.Password) {
		return "", errors.New("password does not match")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(s.cfg.TokenTTL)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})
	return token.SignedString(s.cfg.JWTSigningKey)
}

func (s *AuthService) ValidateToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.cfg.JWTSigningKey, nil
	})
	if err != nil {
		return 0, err
	}
	claim, ok := token.Claims.(*tokenClaims)
	if !ok && !token.Valid {
		return 0, errors.New("invalid token")
	}
	return claim.UserID, nil
}
