package main

import (
	"github.com/sirupsen/logrus"
	"to-do-checklist/internal/app"
	"to-do-checklist/internal/config"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	cfg := config.NewConfig()
	application := app.NewApp(cfg)
	application.Start()
}
