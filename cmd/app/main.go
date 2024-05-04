package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"to-do-checklist/internal/app"
	"to-do-checklist/internal/config"
)

// @title           TODO Lists API
// @version         1.0
// @description     This is an API for TODO Lists Application

// @host localhost:8081
// @BasePath  /

// @securityDefinitions.apikey API Auth Key
// @in header
// @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	gin.SetMode(gin.DebugMode)
	cfg := config.NewConfig()
	application := app.NewApp(cfg)
	application.Start()
}
