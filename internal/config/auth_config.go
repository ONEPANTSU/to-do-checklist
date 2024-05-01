package config

import (
	"github.com/spf13/viper"
	"os"
)

type AuthConfig struct {
	TokenTTL            uint
	PasswordHashingSalt string
	JWTSigningKey       []byte
}

func newAuthConfig() *AuthConfig {
	return &AuthConfig{
		TokenTTL:            viper.GetUint("auth.token_ttl"),
		PasswordHashingSalt: os.Getenv("PASSWORD_HASHING_SALT"),
		JWTSigningKey:       []byte(os.Getenv("JWT_SIGNING_KEY")),
	}
}
