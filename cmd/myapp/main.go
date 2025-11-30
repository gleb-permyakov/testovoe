package main

import (
	"log"
	"main/internal/app"
	"main/internal/config"
)

func init() {
	config.GetConfig()
}

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	application := app.New(cfg)
	err = application.Start(cfg)
	if err != nil {
		log.Printf("error %v", err)
	}

}
