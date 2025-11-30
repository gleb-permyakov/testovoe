package app

import (
	"main/internal/client"
	"main/internal/config"
	"main/internal/handler"
	"main/internal/server"
)

type App struct {
	HTTPsrv *server.HTTPServer
}

func New(cfg *config.Config, client *client.Client) *App {
	r := handler.Route(client)
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
