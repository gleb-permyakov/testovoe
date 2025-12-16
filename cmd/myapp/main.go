package main

import (
	"log"
	"main/internal/app"
	"main/internal/client"
	"main/internal/config"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}

	client := client.New()

	application := app.New(cfg, client)
	err = application.Start(cfg)
	if err != nil {
		log.Printf("error %v\n", err)
	}
}
