package app

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"to-do-checklist/internal/config"
	"to-do-checklist/internal/database"
	httpHandler "to-do-checklist/internal/delivery/http"
	"to-do-checklist/internal/repository/postgres"
	"to-do-checklist/internal/server"
	httpServer "to-do-checklist/internal/server/http"
	"to-do-checklist/internal/service"
)

type App struct {
	server server.Server
	db     *sqlx.DB
}

func NewApp(cfg *config.Config) App {
	db, err := database.Connect(cfg.DB)
	if err != nil {
		logrus.Fatalf("error occurred while db connecting: %s", err)
	}
	repos := postgres.NewPostgresRepository(db)
	services := service.NewService(repos, cfg.Auth)
	handler := httpHandler.NewHandler(services)

	return App{
		httpServer.NewServer(
			cfg.App.Port,
			handler.InitRoutes(),
		),
		db,
	}
}

func (app App) Start() {
	go func() {
		if err := app.server.Run(); err != nil {
			logrus.Fatalf("error occurred while running server: %s", err.Error())
		}
	}()

	logrus.Info("app started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Info("app stopping")

	if err := app.server.Shutdown(context.Background()); err != nil {
		logrus.Fatalf("error occurred while shutting down server: %s", err.Error())
		return
	}
	if err := app.db.Close(); err != nil {
		logrus.Fatalf("error occurred while closing db connection: %s", err.Error())
		return
	}
}
