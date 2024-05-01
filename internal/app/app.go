package app

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"to-do-checklist/internal/config"
	"to-do-checklist/internal/database"
	httpHandler "to-do-checklist/internal/delivery/http"
	"to-do-checklist/internal/repository"
	"to-do-checklist/internal/server"
	httpServer "to-do-checklist/internal/server/http"
	"to-do-checklist/internal/service"
)

type App struct {
	server server.Server
}

func NewApp(cfg *config.Config) App {
	db := database.NewPostgresDB()
	if err := db.Connect(cfg.DB); err != nil {
		logrus.Fatalf("error occurred while db connecting: %s", err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos, cfg.Auth)
	handler := httpHandler.NewHandler(services)

	return App{
		server: httpServer.NewServer(
			cfg.App.Port,
			handler.InitRoutes(),
		),
	}
}

func (app App) Start() {
	if err := app.server.Run(); err != nil {
		logrus.Fatalf("error occurred while running server: %s", err.Error())
	}
}
