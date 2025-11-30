package app

import (
	"main/internal/config"
	"main/internal/handlers"
	"main/internal/server"
)

type App struct {
	HTTPsrv *server.HTTPServer
}

func New(cfg *config.Config) *App {
	r := handlers.Route()
	srv := server.New(cfg, r)
	return &App{
		HTTPsrv: srv,
	}
}

func (a App) Start(cfg *config.Config) error {
	err := a.HTTPsrv.Start(cfg)
	if err != nil {
		return err
	}
	return nil
}
