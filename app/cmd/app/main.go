package main

import (
	"log"
	"news-mono/cmd/internal/app"
	"news-mono/cmd/internal/config"
	"news-mono/cmd/pkg/logging"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logger := logging.GetLogger(cfg.AppConfig.LogLevel)

	a, err := app.NewApp(cfg, &logger)

	if err != nil {
		logger.Fatal(err)
	}

	logger.Println("Running Application")
	a.Run()
}
