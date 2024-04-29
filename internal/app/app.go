package app

import (
	"log"
	httpHandler "to-do-checklist/internal/delivery/http"
	"to-do-checklist/internal/server"
	httpServer "to-do-checklist/internal/server/http"
)

type App struct {
	server server.Server
}

func NewApp(port string) App {
	handler := httpHandler.Handler{}

	return App{
		server: httpServer.NewServer(
			port,
			handler.InitRoutes(),
		),
	}
}

func (app App) Start() {
	if err := app.server.Run(); err != nil {
		log.Fatal("error occurred while running server: %s", err.Error())
	}
}
