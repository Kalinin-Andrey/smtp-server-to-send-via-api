package main

import (
	"context"
	"log"

	"smtp2api/internal/pkg/config"

	commonApp "smtp2api/internal/app"
	"smtp2api/internal/app/restapi"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Get()
	if err != nil {
		log.Fatalln("Can not load the config")
	}
	app := restapi.New(commonApp.New(ctx, *cfg), *cfg)

	if err := app.Run(); err != nil {
		log.Fatalf("Error while application is running: %s", err.Error())
	}
	defer func() {
		if err := app.Stop(); err != nil {
			log.Fatalf("Error while application is stopping: %s", err.Error())
		}
	}()
}
