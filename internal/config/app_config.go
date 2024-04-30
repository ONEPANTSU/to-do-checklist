package config

import "github.com/spf13/viper"

type AppConfig struct {
	Port string
}

func newAppConfig() *AppConfig {
	return &AppConfig{
		Port: viper.GetString("app.port"),
	}
}
