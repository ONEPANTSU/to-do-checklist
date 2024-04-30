package config

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	App *AppConfig
	DB  *DBConfig
}

func initYamlConfig() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error occurred while yaml config initializing: %s", err)
	}
}

func initEnvConfig() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occurred while .env config initializing: %s", err)
	}
}

func NewConfig() *Config {
	initYamlConfig()
	initEnvConfig()
	return &Config{
		App: newAppConfig(),
		DB:  newDBConfig(),
	}
}
